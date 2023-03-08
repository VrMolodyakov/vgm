package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
