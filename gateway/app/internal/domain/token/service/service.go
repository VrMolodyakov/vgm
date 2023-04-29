package service

import (
	"context"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("token-service")
)

type TokenRepo interface {
	Set(ctx context.Context, refreshToken string, userId int, expireAt time.Duration) error
	Get(ctx context.Context, refreshToken string) (int, error)
	Delete(ctx context.Context, refreshToken string) error
}

type tokenService struct {
	storage TokenRepo
}

func NewTokenService(storage TokenRepo) *tokenService {
	return &tokenService{storage: storage}
}

func (t *tokenService) Save(ctx context.Context, refreshToken string, userId int, expireAt time.Duration) error {
	_, span := tracer.Start(ctx, "storage.Save")
	defer span.End()

	if len(refreshToken) == 0 {
		return errors.New("refresh token is empty")
	}
	if userId < 0 {
		return errors.New("user id can't be less than zero")
	}
	return t.storage.Set(ctx, refreshToken, userId, expireAt)
}

func (t *tokenService) Find(ctx context.Context, refreshToken string) (int, error) {
	_, span := tracer.Start(ctx, "storage.Find")
	defer span.End()

	if len(refreshToken) == 0 {
		return -1, errors.New("refresh token is empty")
	}
	return t.storage.Get(ctx, refreshToken)
}

func (t *tokenService) Remove(ctx context.Context, refreshToken string) error {
	_, span := tracer.Start(ctx, "storage.Remove")
	defer span.End()

	if len(refreshToken) == 0 {
		return errors.New("refresh token is empty")
	}
	return t.storage.Delete(ctx, refreshToken)
}
