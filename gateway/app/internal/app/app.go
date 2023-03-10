package app

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/config"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/grpc/v1/client"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/handler/album"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/handler/user"
	userMiddleware "github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/middleware"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/album/service"
	tokenRepo "github.com/VrMolodyakov/vgm/gateway/internal/domain/token/repository"
	tokenService "github.com/VrMolodyakov/vgm/gateway/internal/domain/token/service"
	userRepo "github.com/VrMolodyakov/vgm/gateway/internal/domain/user/repository"
	userService "github.com/VrMolodyakov/vgm/gateway/internal/domain/user/service"
	"github.com/VrMolodyakov/vgm/gateway/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/gateway/pkg/client/redis"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"github.com/VrMolodyakov/vgm/gateway/pkg/token"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	cfg        *config.Config
	httpServer *http.Server
}

func NewApp(cfg *config.Config) *app {
	return &app{cfg: cfg}
}

func (a *app) Run(ctx context.Context) {
	a.startHTTP(ctx)
}

func (a *app) startHTTP(ctx context.Context) error {
	logger := logging.LoggerFromContext(ctx)
	logger.Sugar().Infow("http config:", "port", a.cfg.HTTP.Port, "ip", a.cfg.HTTP.IP)
	logger.Info("HTTP Server initializing")

	pgConfig := postgresql.NewPgConfig(
		a.cfg.Postgres.User,
		a.cfg.Postgres.Password,
		a.cfg.Postgres.IP,
		a.cfg.Postgres.Port,
		a.cfg.Postgres.Database,
		a.cfg.Postgres.PoolSize,
	)

	rdCfg := redis.NewRdConfig(
		a.cfg.Redis.Password,
		a.cfg.Redis.Host,
		a.cfg.Redis.Port,
		a.cfg.Redis.DbNumber,
	)

	rdClient, err := redis.NewClient(ctx, &rdCfg)
	if err != nil {
		logger.Fatal(err.Error())
	}
	pgClient, err := postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	accessKeyPair, refreshKeyPair := a.loadKeyPairs()

	tokenHandler := token.NewTokenHandler(accessKeyPair, refreshKeyPair)

	userRepo := userRepo.NewUserRepo(pgClient)
	tokenRepo := tokenRepo.NewTokenRepo(rdClient)
	userService := userService.NewUserService(userRepo)
	tokenService := tokenService.NewTokenService(tokenRepo)

	userHandler := user.NewUserHandler(userService, tokenHandler, tokenService, a.cfg.KeyPairs.AccessTtl, a.cfg.KeyPairs.RefreshTtl)

	userAuth := userMiddleware.NewAuthMiddleware(userService, tokenService, tokenHandler)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	grpcClient := client.NewMusicClient("music:30000")
	grpcClient.Start()
	albumService := service.NewAlbumService(grpcClient)
	albumHandler := album.NewAlbumHandler(albumService)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/register", userHandler.SignUpUser)
		r.Post("/login", userHandler.SignInUser)
		r.Get("/refresh", userHandler.RefreshAccessToken)
		r.Group(func(r chi.Router) {
			r.Use(userAuth.Auth)
			r.Get("/logout", userHandler.Logout)
		})
	})

	router.Route("/music", func(r chi.Router) {
		r.Use(userAuth.Auth)
		r.Post("/create", albumHandler.CreateAlbum)
		r.Get("/albums", albumHandler.FindAllAlbums)
	})

	addr := fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port)
	fmt.Println(addr)

	a.httpServer = &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}
	if err = a.httpServer.ListenAndServe(); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err.Error())
		}
	}
	err = a.httpServer.Shutdown(ctx)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return err
}

func (a *app) loadKeyPairs() (token.KeyPair, token.KeyPair) {
	aprk, err := base64.StdEncoding.DecodeString(a.cfg.KeyPairs.AccessPrivate)
	if err != nil {
		log.Fatal(err)
	}
	apbk, err := base64.StdEncoding.DecodeString(a.cfg.KeyPairs.AccessPublic)
	if err != nil {
		log.Fatal(err)
	}
	rprk, err := base64.StdEncoding.DecodeString(a.cfg.KeyPairs.RefreshPrivate)
	if err != nil {
		log.Fatal(err)
	}
	rpbk, err := base64.StdEncoding.DecodeString(a.cfg.KeyPairs.RefreshPublic)
	if err != nil {
		log.Fatal(err)
	}
	apair := token.KeyPair{PrivateKey: aprk, PublicKey: apbk}
	rpair := token.KeyPair{PrivateKey: rprk, PublicKey: rpbk}
	return apair, rpair
}
