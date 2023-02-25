package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
)

type UserRepo interface {
	Create(ctx context.Context, user model.User) error
	GetOne(ctx context.Context, username string) (model.User, error)
	Delete(ctx context.Context, username string) error
	Update(ctx context.Context, user model.User) error
}

type userService struct {
	userRepo UserRepo
}

func NewUserService(repo UserRepo) *userService {
	return &userService{
		userRepo: repo,
	}
}

func (u *userService) Create(ctx context.Context, user model.User) error {
	if !user.IsValid() {
		return errors.New("user data must not be empty")
	}
	return u.userRepo.Create(ctx, user)
}

func (u *userService) GetOne(ctx context.Context, username string) (model.User, error) {
	if username == "" {
		return model.User{}, errors.New("username is empty")
	}
	return u.userRepo.GetOne(ctx, username)
}

func (u *userService) Delete(ctx context.Context, username string) error {
	if username == "" {
		return errors.New("username is empty")
	}
	return u.userRepo.Delete(ctx, username)
}

func (u *userService) Update(ctx context.Context, user model.User) error {
	if !user.IsValid() {
		return errors.New("user data must not be empty")
	}
	return u.userRepo.Update(ctx, user)
}
