package app

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/config"
	userRepo "github.com/VrMolodyakov/vgm/gateway/internal/domain/user/repository"
	"github.com/VrMolodyakov/vgm/gateway/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/gateway/pkg/client/redis"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"github.com/VrMolodyakov/vgm/gateway/pkg/token"
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

	pgConfig := postgresql.NewPgConfig(
		a.cfg.Postgres.User,
		a.cfg.Postgres.Password,
		a.cfg.Postgres.IP,
		a.cfg.Postgres.Port,
		a.cfg.Postgres.Database,
		a.cfg.Postgres.PoolSize,
	)
	rdCfg := redis.NewRdConfig(a.cfg.Redis.Password, a.cfg.Redis.Host, a.cfg.Redis.Port, a.cfg.Redis.DbNumber)
	rdClient, err := redis.NewClient(ctx, &rdCfg)
	if err != nil {
		logger.Fatal(err.Error())
	}
	pgClient, err := postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	userRepo := userRepo.NewUserRepo(pgClient)

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

func (a *app) loadTokens() (token.TokenPair, token.TokenPair) {
	aprk, err := base64.StdEncoding.DecodeString(a.cfg.TokenPairs.AccessPrivate)
	if err != nil {
		log.Fatal(err)
	}
	apbk, err := base64.StdEncoding.DecodeString(a.cfg.TokenPairs.AccessPublic)
	if err != nil {
		log.Fatal(err)
	}
	rprk, err := base64.StdEncoding.DecodeString(a.cfg.TokenPairs.RefreshPrivate)
	if err != nil {
		log.Fatal(err)
	}
	rpbk, err := base64.StdEncoding.DecodeString(a.cfg.TokenPairs.RefreshPublic)
	if err != nil {
		log.Fatal(err)
	}
	apair := token.TokenPair{PrivateKey: aprk, PublicKey: apbk}
	rpair := token.TokenPair{PrivateKey: rprk, PublicKey: rpbk}
	return apair, rpair
}
