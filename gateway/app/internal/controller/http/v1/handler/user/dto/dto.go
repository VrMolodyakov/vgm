package dto

import (
	validate "github.com/go-playground/validator/v10"
)

var (
	validator = validate.New()
)

type SignUpRequest struct {
	Username string `json:"username" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role"`
}

type SignInRequest struct {
	Username string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponse struct {
	UserID int `json:"user_id"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	LoggedIn     string `json:"logged_in"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	LoggedIn     string `json:"logged_in"`
}

type ReqError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func (r *SignUpRequest) Validate() []ReqError {
	var errors []ReqError
	err := validator.Struct(r)
	if err != nil {
		for _, err := range err.(validate.ValidationErrors) {
			var re ReqError
			re.Field = err.Field()
			re.Tag = err.Tag()
			re.Param = err.Param()
			errors = append(errors, re)
		}
	}
	return errors
}
