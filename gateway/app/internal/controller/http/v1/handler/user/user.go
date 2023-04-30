package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/handler/user/dto"
	emodel "github.com/VrMolodyakov/vgm/gateway/internal/domain/email/model"
	umodel "github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/email/templates"

	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/hashing"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("user-http")
)

const (
	link        string        = "http://localhost:3000/home"
	sendTimeout time.Duration = 30 * time.Second
)

type userHandler struct {
	user         UserService
	tokenHandler TokenHandler
	tokenService TokenService
	email        EmailClient
	accessTtl    int
	refreshTtl   int
}

func NewUserHandler(
	user UserService,
	tokenHandler TokenHandler,
	tokenService TokenService,
	email EmailClient,
	accessTtl int,
	refreshTtl int) *userHandler {

	return &userHandler{
		user:         user,
		tokenHandler: tokenHandler,
		tokenService: tokenService,
		accessTtl:    accessTtl,
		email:        email,
		refreshTtl:   refreshTtl,
	}
}

func (u *userHandler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var req dto.SignUpRequest
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashing.HashPassword(req.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	user := umodel.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		Role:     req.Role,
	}
	userID, err := u.user.Create(r.Context(), user)
	if err != nil {
		if _, ok := errors.IsInternal(err); ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go u.sendEmail(user)

	response := dto.UserResponse{UserID: userID}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (u *userHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()

	var req dto.SignInRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}
	user, err := u.user.GetByUsername(context.Background(), req.Username)
	if err != nil {
		if _, ok := errors.IsInternal(err); ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = hashing.ComparePassword(user.Password, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accessToken, err := u.tokenHandler.CreateAccessToken(time.Duration(u.accessTtl)*time.Minute, user.Id, user.Role)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	refreshToken, err := u.tokenHandler.CreateRefreshToken(time.Duration(u.refreshTtl)*time.Minute, user.Id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = u.tokenService.Save(r.Context(), refreshToken, user.Id, time.Duration(u.refreshTtl)*time.Minute)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	accessCookie := http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   u.accessTtl * 60,
		Secure:   false,
		HttpOnly: true,
	}

	refreshCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   u.refreshTtl * 60,
		Secure:   false,
		HttpOnly: true,
	}

	loginCookie := http.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   u.accessTtl * 60,
		Secure:   false,
		HttpOnly: false,
	}

	response := dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		LoggedIn:     "true",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &accessCookie)
	http.SetCookie(w, &refreshCookie)
	http.SetCookie(w, &loginCookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (u *userHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()

	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	refreshToken := refreshTokenCookie.Value
	err = u.tokenHandler.ValidateRefreshToken(refreshToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	userId, err := u.tokenService.Find(r.Context(), refreshToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	user, err := u.user.GetByID(r.Context(), userId)
	if err != nil {
		if _, ok := errors.IsInternal(err); ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	accessToken, err := u.tokenHandler.CreateAccessToken(time.Duration(u.accessTtl)*time.Minute, userId, user.Role)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	accessCookie := http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   u.accessTtl * 60,
		Secure:   false,
		HttpOnly: true,
	}
	loginCookie := http.Cookie{
		Name:     "ogged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   u.accessTtl * 60,
		Secure:   false,
		HttpOnly: false,
	}

	response := dto.RefreshTokenResponse{
		AccessToken: accessToken,
		LoggedIn:    "true",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &accessCookie)
	http.SetCookie(w, &loginCookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (u *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()
	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	refreshToken := refreshTokenCookie.Value
	err = u.tokenHandler.ValidateRefreshToken(refreshToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	err = u.tokenService.Remove(r.Context(), refreshToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	accessCookie := http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   u.accessTtl * 60,
		Secure:   false,
		HttpOnly: true,
	}

	refreshCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		MaxAge:   u.refreshTtl * 60,
		Secure:   false,
		HttpOnly: true,
	}

	loginCookie := http.Cookie{
		Name:     "logged_in",
		Value:    "",
		Path:     "/",
		MaxAge:   u.accessTtl * 60,
		Secure:   false,
		HttpOnly: false,
	}

	response := dto.TokenResponse{
		AccessToken:  "",
		RefreshToken: "",
		LoggedIn:     "",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &accessCookie)
	http.SetCookie(w, &refreshCookie)
	http.SetCookie(w, &loginCookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

//TODO: change To
func (u *userHandler) sendEmail(user umodel.User) {
	ctx, cancel := context.WithTimeout(context.Background(), sendTimeout)
	_, span := tracer.Start(ctx, "send email")
	defer span.End()
	defer cancel()
	logger := logging.GetLogger()
	t := templates.NewTemplate(link)
	content, err := t.Greeting(user.Username, link)
	if err != nil {
		return
	}
	email := emodel.Email{
		Subject: "Greeting",
		Content: content,
		To:      []string{"vrmolodyakov@mail.ru"},
	}
	err = u.email.Send(ctx, email)
	if err != nil {
		logger.Error(err.Error())
	}
}
