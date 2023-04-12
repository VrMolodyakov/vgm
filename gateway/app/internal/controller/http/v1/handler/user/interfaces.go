package user

import (
	"context"
	"time"

	emodel "github.com/VrMolodyakov/vgm/gateway/internal/domain/email/model"
	umodel "github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
)

type UserService interface {
	Create(ctx context.Context, user umodel.User) (int, error)
	GetByUsername(ctx context.Context, username string) (umodel.User, error)
	GetByID(ctx context.Context, ID int) (umodel.User, error)
}

type TokenHandler interface {
	CreateAccessToken(ttl time.Duration, payload interface{}, role string) (string, error)
	CreateRefreshToken(ttl time.Duration, payload interface{}) (string, error)
	ValidateRefreshToken(token string) error
}

type TokenService interface {
	Save(ctx context.Context, refreshToken string, userId int, expireAt time.Duration) error
	Find(ctx context.Context, refreshToken string) (int, error)
	Remove(ctx context.Context, refreshToken string) error
}

type EmailClient interface {
	Send(ctx context.Context, email emodel.Email) error
}
