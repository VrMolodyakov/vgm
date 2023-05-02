package dto

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
