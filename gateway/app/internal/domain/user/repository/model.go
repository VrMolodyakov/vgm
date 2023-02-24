package repository

import (
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	mapper "github.com/worldline-go/struct2"
)

type userStorage struct {
	Username string    `struct:"user_name"`
	Email    string    `struct:"user_mail"`
	Password string    `struct:"user_password"`
	CreateAt time.Time `struct:"create_at"`
}

func toStorage(user model.User) userStorage {
	return userStorage{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		CreateAt: time.Now(),
	}
}

func toStorageMap(user model.User) map[string]interface{} {
	storage := toStorage(user)
	return (&mapper.Decoder{}).Map(storage)
}
