package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
)

type UserService interface {
	GetById(ctx context.Context, id int) (model.User, error)
}

type TokenHandler interface {
	ValidateAccessToken(token string) (interface{}, error)
}

type TokenService interface {
	Find(refreshToken string) (int, error)
}

type authMiddleware struct {
	userService  UserService
	tokenHandler TokenHandler
	tokenService TokenService
}

func NewAuthMiddleware(
	userService UserService,
	tokenService TokenService,
	tokenHandler TokenHandler) *authMiddleware {
	return &authMiddleware{userService: userService,
		tokenService: tokenService,
		tokenHandler: tokenHandler}
}

func (a *authMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accessToken string
		coockie, err := r.Cookie("access_token")
		authHeader := r.Header.Get("Authorization")
		fields := strings.Fields(authHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = coockie.Value
		}
		if accessToken == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		sub, err := a.tokenHandler.ValidateAccessToken(accessToken)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		userId := sub.(float64)
		user, err := a.userService.GetById(r.Context(), int(userId))
		if err != nil {
			if _, ok := errors.IsInternal(err); ok {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
