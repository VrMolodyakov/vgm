package middleware

import (
	"context"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
)

type UserService interface {
	GetByID(ctx context.Context, ID int) (model.User, error)
}

type TokenHandler interface {
	ValidateAccessToken(token string) (interface{}, error)
}

type TokenService interface {
	Find(ctx context.Context, refreshToken string) (int, error)
}

type authMiddleware struct {
	userService  UserService
	tokenHandler TokenHandler
	tokenService TokenService
}
