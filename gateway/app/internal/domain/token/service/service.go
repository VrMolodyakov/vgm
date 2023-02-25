package service

import (
	"context"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
)

type TokenStorage interface {
	Set(refreshToken string, userId int, expireAt time.Duration) error
	Get(refreshToken string) (int, error)
	Delete(refreshToken string) error
}

type tokenService struct {
	storage TokenStorage
}

func NewTokenService(storage TokenStorage) *tokenService {
	return &tokenService{storage: storage}
}

func (t *tokenService) Save(ctx context.Context, refreshToken string, userId int, expireAt time.Duration) error {
	if len(refreshToken) == 0 {
		return errors.New("refresh token is empty")
	}
	if userId < 0 {
		return errors.New("user id can't be less than zero")
	}
	return t.storage.Set(refreshToken, userId, expireAt)
}

func (t *tokenService) Find(ctx context.Context, refreshToken string) (int, error) {
	if len(refreshToken) == 0 {
		return -1, errors.New("refresh token is empty")
	}
	return t.storage.Get(refreshToken)
}

func (t *tokenService) Remove(ctx context.Context, refreshToken string) error {
	if len(refreshToken) == 0 {
		return errors.New("refresh token is empty")
	}
	return t.storage.Delete(refreshToken)
}
