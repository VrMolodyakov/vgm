package user

import (
	"context"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
)

type UserService interface {
	Create(ctx context.Context, user model.User) error
	GetOne(ctx context.Context, username string) (model.User, error)
	Delete(ctx context.Context, username string) error
	Update(ctx context.Context, user model.User) error
}

type TokenHandler interface {
	CreateAccessToken(ttl time.Duration, payload interface{}) (string, error)
	CreateRefreshToken(ttl time.Duration, payload interface{}) (string, error)
	ValidateRefreshToken(token string) error
}

type TokenService interface {
	Save(refreshToken string, userId int, expireAt time.Duration) error
	Find(refreshToken string) (int, error)
	Remove(refreshToken string) error
}
