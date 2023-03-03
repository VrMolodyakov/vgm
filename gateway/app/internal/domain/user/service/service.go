package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
)

type UserRepo interface {
	Create(ctx context.Context, user model.User) (int, error)
	Delete(ctx context.Context, username string) error
	Update(ctx context.Context, user model.User) error
	GetByUsername(ctx context.Context, username string) (model.User, error)
	GetByID(ctx context.Context, ID string) (model.User, error)
}

type userService struct {
	userRepo UserRepo
}

func NewUserService(repo UserRepo) *userService {
	return &userService{
		userRepo: repo,
	}
}

func (u *userService) Create(ctx context.Context, user model.User) (int, error) {
	if !user.IsValid() {
		return -1, errors.New("user data must not be empty")
	}
	return u.userRepo.Create(ctx, user)
}

func (u *userService) GetByUsername(ctx context.Context, username string) (model.User, error) {
	if username == "" {
		return model.User{}, errors.New("username is empty")
	}
	return u.userRepo.GetByUsername(ctx, username)
}

func (u *userService) GetByID(ctx context.Context, ID string) (model.User, error) {
	if ID == "" {
		return model.User{}, errors.New("ID is empty")
	}
	return u.userRepo.GetByID(ctx, ID)
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
