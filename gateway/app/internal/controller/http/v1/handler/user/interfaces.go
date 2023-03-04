package user

import (
	"context"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
)

type UserService interface {
	Create(ctx context.Context, user model.User) (int, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
}

type TokenHandler interface {
	CreateAccessToken(ttl time.Duration, payload interface{}) (string, error)
	CreateRefreshToken(ttl time.Duration, payload interface{}) (string, error)
	ValidateRefreshToken(token string) error
}

type TokenService interface {
	Save(ctx context.Context, refreshToken string, userId int, expireAt time.Duration) error
	Find(ctx context.Context, refreshToken string) (int, error)
	Remove(ctx context.Context, refreshToken string) error
}
