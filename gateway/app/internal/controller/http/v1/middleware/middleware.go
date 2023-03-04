package middleware

import (
	"net/http"
	"strings"

	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
)

func NewAuthMiddleware(
	userService UserService,
	tokenService TokenService,
	tokenHandler TokenHandler) *authMiddleware {
	return &authMiddleware{
		userService:  userService,
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
		_, err = a.userService.GetByID(r.Context(), int(userId))
		if err != nil {
			if _, ok := errors.IsInternal(err); ok {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
