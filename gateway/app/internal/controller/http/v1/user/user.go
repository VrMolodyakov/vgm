package auth

import (
	"context"
	"net/http"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
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
	// var request UserRequest

	// err := ctx.ShouldBindJSON(&request)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, u.logger, errs.New(errs.Validation, errs.Code("incorrect data format")))
	// 	return
	// }
	// hashedPassword, err := hashing.HashPassword(request.Password)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, u.logger, err)
	// 	return
	// }
	// user, err := u.userService.Create(ctx, request.Username, hashedPassword)
	// if err != nil {
	// 	errs.HTTPErrorResponse(ctx, u.logger, err)
	// 	return
	// }
	// response := ResponseFromEntity(user)
	// ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": response}})

}
