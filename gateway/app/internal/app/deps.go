package app

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/config"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/grpc/v1/client"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/grpc/v1/client/email"
	musicClient "github.com/VrMolodyakov/vgm/gateway/internal/controller/grpc/v1/client/music"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/middleware"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/music"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/user"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/music/service"
	tokenRepo "github.com/VrMolodyakov/vgm/gateway/internal/domain/token/repository"
	tokenService "github.com/VrMolodyakov/vgm/gateway/internal/domain/token/service"
	userRepo "github.com/VrMolodyakov/vgm/gateway/internal/domain/user/repository"
	userService "github.com/VrMolodyakov/vgm/gateway/internal/domain/user/service"
	"github.com/VrMolodyakov/vgm/gateway/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/gateway/pkg/client/redis"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"github.com/VrMolodyakov/vgm/gateway/pkg/token"
)

type Deps struct {
	userServer  *http.Server
	musicServer *http.Server
}

func (d *Deps) Setup(ctx context.Context, cfg config.Config) {
	logger := logging.LoggerFromContext(ctx)
	logger.Info("Setup...")

	pgConfig := postgresql.NewPgConfig(
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.IP,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
		cfg.Postgres.PoolSize,
	)

	rdCfg := redis.NewRdConfig(
		cfg.Redis.Password,
		cfg.Redis.Host,
		cfg.Redis.Port,
		cfg.Redis.DbNumber,
	)

	rdClient, err := redis.NewClient(ctx, &rdCfg)
	if err != nil {
		logger.Fatal(err.Error())
	}

	pgClient, err := postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	accessKeyPair, refreshKeyPair := d.loadKeyPairs(cfg.KeyPairs)

	tokenManager := token.NewTokenManager(accessKeyPair, refreshKeyPair)

	musicAddress := fmt.Sprintf("%s:%d", cfg.MusicGRPC.HostName, cfg.MusicGRPC.Port)
	emailAddress := fmt.Sprintf("%s:%d", cfg.EmailGRPC.HostName, cfg.EmailGRPC.Port)

	emailCerts := d.loadEmailClientCert(cfg.EmailClientCert)
	musicCerts := d.loadMusicClientCert(cfg.MusicClientCert)

	grpcMusicClient := musicClient.NewMusicClient(musicAddress)
	grpcEmailClient := email.NewEmailClient(emailAddress)
	grpcMusicClient.StartWithTLS(musicCerts)
	grpcEmailClient.StartWithTLS(emailCerts)

	userRepo := userRepo.NewUserRepo(pgClient)
	tokenRepo := tokenRepo.NewTokenRepo(rdClient)
	userService := userService.NewUserService(userRepo)
	tokenService := tokenService.NewTokenService(tokenRepo)
	albumService := service.NewAlbumService(grpcMusicClient)

	cors := d.setupCORS(cfg.CORS)
	auth := middleware.NewAuthMiddleware(userService, tokenService, tokenManager)

	d.userServer = user.NewServer(
		userService,
		tokenManager,
		tokenService,
		grpcEmailClient,
		cors,
		auth,
		cfg.KeyPairs,
		cfg.UserServer,
	)

	d.musicServer = music.NewServer(
		cfg.MusicServer,
		cors,
		auth,
		albumService,
	)
}

func (d *Deps) loadKeyPairs(cfg config.KeyPairs) (token.KeyPair, token.KeyPair) {
	aprk, err := base64.StdEncoding.DecodeString(cfg.AccessPrivate)
	if err != nil {
		log.Fatal(err)
	}
	apbk, err := base64.StdEncoding.DecodeString(cfg.AccessPublic)
	if err != nil {
		log.Fatal(err)
	}
	rprk, err := base64.StdEncoding.DecodeString(cfg.RefreshPrivate)
	if err != nil {
		log.Fatal(err)
	}
	rpbk, err := base64.StdEncoding.DecodeString(cfg.RefreshPublic)
	if err != nil {
		log.Fatal(err)
	}
	apair := token.KeyPair{PrivateKey: aprk, PublicKey: apbk}
	rpair := token.KeyPair{PrivateKey: rprk, PublicKey: rpbk}
	return apair, rpair
}

func (d *Deps) loadMusicClientCert(cfg config.MusicClientCert) client.ClientCerts {
	return client.NewClientCerts(
		cfg.EnableTLS,
		cfg.ClientCertFile,
		cfg.ClientKeyFile,
		cfg.ClientCACertFile,
	)
}

func (d *Deps) loadEmailClientCert(cfg config.EmailClientCert) client.ClientCerts {
	return client.NewClientCerts(
		cfg.EnableTLS,
		cfg.ClientCertFile,
		cfg.ClientKeyFile,
		cfg.ClientCACertFile,
	)
}

func (d *Deps) setupCORS(cfg config.CORS) middleware.Cors {
	origins := strings.Join(cfg.AllowedOrigins[:], ", ")
	headers := strings.Join(cfg.AllowedHeaders[:], ", ")
	methods := strings.Join(cfg.AllowedMethods[:], ", ")

	return middleware.NewCors(origins, headers, methods)

}
