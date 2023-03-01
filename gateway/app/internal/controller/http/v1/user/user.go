package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/user/dto"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/hashing"
)

type userHandler struct {
	user         UserService
	tokenHandler TokenHandler
	tokenService TokenService
	accessTtl    int
	refreshTtl   int
}

func NewUserHandler(
	user UserService,
	tokenHandler TokenHandler,
	accessTtl int,
	refreshTtl int) *userHandler {

	return &userHandler{
		user:         user,
		tokenHandler: tokenHandler,
	}
}

func (u *userHandler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRequest

	defer r.Body.Close()

	if err := json.NewEncoder(w).Encode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}
	hashedPassword, err := hashing.HashPassword(req.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	user := model.NewUser(req.Username, hashedPassword, req.Email, req.Role)
	err = u.user.Create(r.Context(), user)
	if err != nil {
		if _, ok := errors.IsInternal(err); ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *userHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRequest
	defer r.Body.Close()

	if err := json.NewEncoder(w).Encode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}
	user, err := u.user.GetOne(context.Background(), req.Username)
	if err != nil {
		if _, ok := errors.IsInternal(err); ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = hashing.ComparePassword(user.Password, req.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	accessToken, err := u.tokenHandler.CreateAccessToken(time.Duration(u.accessTtl)*time.Minute, user.Id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	refreshToken, err := u.tokenHandler.CreateRefreshToken(time.Duration(u.refreshTtl)*time.Minute, user.Id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = u.tokenService.Save(refreshToken, user.Id, time.Duration(u.refreshTtl)*time.Minute)
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
	http.SetCookie(w, &accessCookie)

	refreshCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   u.refreshTtl * 60,
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(w, &refreshCookie)

	loginCookie := http.Cookie{
		Name:     "ogged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   u.accessTtl * 60,
		Secure:   false,
		HttpOnly: false,
	}
	http.SetCookie(w, &loginCookie)

	w.WriteHeader(http.StatusOK)
}

func (u *userHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
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
	// refreshToken, err := ctx.Cookie("refresh_token")
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, errs.New(errs.Unauthorized, err))
	// 	return
	// }
	// a.logger.Infof("get refreshToken from cookie = %v", refreshToken)
	// err = a.tokenHandler.ValidateRefreshToken(refreshToken)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, errs.New(errs.Unauthorized, err))
	// 	return
	// }
	// userId, err := a.tokenService.Find(refreshToken)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, err)
	// 	return
	// }
	// accessToken, err := a.tokenHandler.CreateAccessToken(time.Duration(a.refreshTtl)*time.Minute, userId)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, errs.New(errs.Internal, err))
	// 	return
	// }
	// ctx.SetCookie("access_token", accessToken, a.accessTtl*60, "/", a.host, false, true)
	// ctx.SetCookie("logged_in", "true", a.accessTtl*60, "/", a.host, false, false)

	// ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}
