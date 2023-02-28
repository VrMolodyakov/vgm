package model

import (
	"time"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
	Role     string
	CreateAt time.Time
}

func (u *User) IsValid() bool {
	return u.Username != "" && u.Email != "" && u.Password != "" && u.Role != ""
}

func NewUser(username, password, email, role string) User {
	return User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}
}
