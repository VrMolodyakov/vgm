package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/VrMolodyakov/vgm/gateway/internal/config"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"github.com/go-chi/chi"
)

type app struct {
	cfg        *config.Config
	httpServer *http.Server
}

func NewApp(cfg *config.Config) *app {
	return &app{cfg: cfg}
}

func (a *app) startHTTP(ctx context.Context) error {
	logger := logging.LoggerFromContext(ctx)
	logger.Sugar().Infow("http config:", "port", a.cfg.HTTP.Port, "ip", a.cfg.HTTP.IP)
	logger.Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		logger.Info(err.Error())
		logger.Fatal("failed to create listener")
	}

	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// router.Get()
	// router := rg.Group("/auth")
	// router.POST("/register", a.authHandler.SignUpUser)
	// router.POST("/login", a.authHandler.SignInUser)
	// router.GET("/refresh", a.authHandler.RefreshAccessToken)
	// router.GET("/logout", a.authMiddleware.Auth(), a.authHandler.Logout)
	// handler := c.Handler(a.router)

	// router.Route("/auth", func(r chi.Router) {
	// 	r.Get("/", getArticle)
	// 	r.Put("/", updateArticle)
	// 	r.Delete("/", deleteArticle)
	//   })

	a.httpServer = &http.Server{
		Handler:      router,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err.Error())
		}
	}
	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		logger.Fatal(err.Error())
	}
	return err
}
