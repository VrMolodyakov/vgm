package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/VrMolodyakov/vgm/gateway/internal/config"
	"github.com/VrMolodyakov/vgm/gateway/pkg/jaeger"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
)

type app struct {
	cfg  *config.Config
	deps Deps
}

func New() *app {
	logging.Init("info", "log.txt")
	logger := logging.GetLogger()
	logger.Info("New app")

	return &app{}
}

func (a *app) Close(ctx context.Context) {
	a.deps.Close(ctx)
}

func (a *app) Setup(ctx context.Context) error {
	return a.deps.Setup(ctx, a.cfg)
}

func (a *app) ReadConfig() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	a.cfg = cfg
	return nil
}

func (a *app) InitTracer() error {
	err := jaeger.SetGlobalTracer(a.cfg.Jaeger.ServiceName, a.cfg.Jaeger.Address, a.cfg.Jaeger.Port)
	if err != nil {
		return err
	}
	return nil
}

func (a *app) Start() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		logger := logging.GetLogger()
		logger.Sugar().Info("music server started on ", " addr ", a.cfg.MusicServer.Port)
		if err := a.deps.musicServer.ListenAndServe(); err != nil {
			switch {
			case errors.Is(err, http.ErrServerClosed):
				logger.Warn("server shutdown")
			default:
				logger.Fatal(err.Error())
			}
		}
		err := a.deps.musicServer.Shutdown(ctx)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}()

	go func() {
		logger := logging.GetLogger()
		logger.Sugar().Info("user server started on ", " addr ", a.cfg.UserServer.Port)
		if err := a.deps.userServer.ListenAndServe(); err != nil {
			switch {
			case errors.Is(err, http.ErrServerClosed):
				logger.Warn("server shutdown")
			default:
				logger.Fatal(err.Error())
			}
		}
		err := a.deps.userServer.Shutdown(ctx)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}()

	<-ctx.Done()
}

func (a *app) LogConfig() {
	logger := logging.GetLogger()
	logger.Sugar().Infow(
		"psql config",
		"database", a.cfg.Postgres.Database,
		"IP", a.cfg.Postgres.IP,
		"Password", a.cfg.Postgres.Password,
		"Pool size", a.cfg.Postgres.PoolSize,
		"Port", a.cfg.Postgres.Port,
		"User", a.cfg.Postgres.User,
	)

	logger.Sugar().Infow(
		"redis config",
		"Host", a.cfg.Redis.Host,
		"Password", a.cfg.Redis.Password,
		"Port", a.cfg.Redis.Port,
		"Db number", a.cfg.Redis.DbNumber,
	)

	logger.Sugar().Infow(
		"jaeger",
		"Address", a.cfg.Jaeger.Address,
		"Port", a.cfg.Jaeger.Port,
	)

	logger.Sugar().Infow(
		"keys config",
		"Access private key", a.cfg.KeyPairs.AccessPrivate,
		"Access public key", a.cfg.KeyPairs.AccessPublic,
		"Refresh private key", a.cfg.KeyPairs.RefreshPrivate,
		"Refresh public key", a.cfg.KeyPairs.RefreshPublic,
		"Access TTL", a.cfg.KeyPairs.AccessTtl,
		"Refresh TTL", a.cfg.KeyPairs.RefreshTtl,
	)
}
