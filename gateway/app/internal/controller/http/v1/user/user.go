package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/user/dto"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/hashing"
)

type UserService interface {
	Create(ctx context.Context, user model.User) error
	GetOne(ctx context.Context, username string) (model.User, error)
	Delete(ctx context.Context, username string) error
	Update(ctx context.Context, user model.User) error
}

type userHandler struct {
	user UserService
}

func NewUserHandler(user UserService) *userHandler {
	return &userHandler{
		user: user,
	}
}

func (u *userHandler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRequest

	defer r.Body.Close()

	if err := json.NewEncoder(w).Encode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), 400)
		return
	}
	hashedPassword, err := hashing.HashPassword(req.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
	user := model.NewUser(req.Username, hashedPassword, req.Email, req.Role)
	err = u.user.Create(r.Context(), user)
	if err != nil {
		if _, ok := errors.ISInternal(err); ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), 401)
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *userHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRequest
	defer r.Body.Close()

	if err := json.NewEncoder(w).Encode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), 400)
		return
	}
	user, err := u.user.GetOne(context.Background(), req.Username)
	if err != nil {
		if _, ok := errors.ISInternal(err); ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		}
		http.Error(w, http.StatusText(http.StatusBadRequest), 400)
	}
	fmt.Println(user)
	// var request UserRequest

	// err := ctx.ShouldBindJSON(&request)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, err)
	// 	return
	// }

	// user, err := a.userService.Get(ctx, request.Username)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, err)
	// 	return
	// }
	// a.logger.Debugf("find user = %v", user)
	// err = hashing.ComparePassword(user.Password, request.Password)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, errs.New(errs.Validation, errs.Code("wrong password"), errs.Parameter("password"), err))
	// 	return
	// }

	// accessToken, err := a.tokenHandler.CreateAccessToken(time.Duration(a.accessTtl)*time.Minute, user.Id)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, errs.New(errs.Internal, err))
	// 	return
	// }
	// refreshToken, err := a.tokenHandler.CreateRefreshToken(time.Duration(a.refreshTtl)*time.Minute, user.Id)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, errs.New(errs.Internal, err))
	// 	return
	// }

	// err = a.tokenService.Save(refreshToken, user.Id, time.Duration(a.refreshTtl)*time.Minute)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, a.logger, err)
	// 	return
	// }

	// ctx.SetCookie("access_token", accessToken, a.accessTtl*60, "/", a.host, false, true)
	// ctx.SetCookie("refresh_token", refreshToken, a.refreshTtl*60, "/", a.host, false, true)
	// ctx.SetCookie("logged_in", "true", a.accessTtl*60, "/", a.host, false, false)

	// ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}
