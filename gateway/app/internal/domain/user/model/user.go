package model

import "time"

type User struct {
	Id       int
	Username string
	Email    string
	Password string
	CreateAt time.Time
}

func (u *User) IsValid() bool {
	return u.Username != "" && u.Email != "" && u.Password != ""
}
