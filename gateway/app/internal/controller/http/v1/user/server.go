package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/config"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/middleware"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewServer(
	user UserService,
	tokenHandler TokenManager,
	tokenService TokenService,
	email EmailClient,
	cors middleware.Cors,
	auth *middleware.AuthMiddleware,
	tokenCfg config.KeyPairs,
	serverCfg config.UserServer) *http.Server {

	handler := NewUserHandler(
		user,
		tokenHandler,
		tokenService,
		email,
		tokenCfg.AccessTtl,
		tokenCfg.RefreshTtl,
	)

	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)
	router.Use(cors.CORS)
	router.Use(middleware.DurationMiddleware)
	router.Use(chiMiddleware.Recoverer)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.SignUpUser)
		r.Post("/login", handler.SignInUser)
		r.Get("/refresh", handler.RefreshAccessToken)
		r.Group(func(r chi.Router) {
			r.Use(auth.Auth)
			r.Get("/logout", handler.Logout)
		})
	})

	addr := fmt.Sprintf("%s:%d", serverCfg.IP, serverCfg.Port)
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: time.Duration(serverCfg.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(serverCfg.ReadTimeout) * time.Second,
	}
}
