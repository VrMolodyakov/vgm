package repository

import (
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	mapper "github.com/worldline-go/struct2"
)

const (
	fields = 4
)

type userStorage struct {
	Username string    `struct:"user_name"`
	Email    string    `struct:"user_email"`
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

func toUpdateStorageMap(m model.User) map[string]interface{} {

	storageMap := make(map[string]interface{}, fields)

	if m.Username != "" {
		storageMap["user_name"] = m.Username
	}
	if m.Email != "" {
		storageMap["user_email"] = m.Email
	}
	if m.Password != "" {
		storageMap["user_password"] = m.Password
	}
	if !m.CreateAt.IsZero() {
		storageMap["create_at"] = m.CreateAt
	}
	return storageMap
}
