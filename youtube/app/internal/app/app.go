package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/VrMolodyakov/vgm/youtube/internal/config"
	"github.com/VrMolodyakov/vgm/youtube/pkg/jaeger"
	"github.com/VrMolodyakov/vgm/youtube/pkg/logging"
)

type app struct {
	cfg    *config.Config
	logger logging.Logger
	deps   Deps
}

func New() *app {
	return &app{}
}

func (a *app) Setup(ctx context.Context) error {
	return a.deps.Setup(ctx, a.cfg, a.logger)
}

func (a *app) InitLogger() {
	loggerCfg := logging.NewLogerConfig(
		a.cfg.Logger.DisableCaller,
		a.cfg.Logger.Development,
		a.cfg.Logger.DisableStacktrace,
		a.cfg.Logger.Encoding,
		a.cfg.Logger.Level,
	)
	a.logger = logging.NewLogger(loggerCfg)
	a.logger.InitLogger()
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

func (a *app) Start(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, os.Kill, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		a.logger.Info("youtube server started on ", " addr ", a.cfg.YoutubeServer.Port)
		if err := a.deps.youtubeServer.ListenAndServe(); err != nil {
			switch {
			case errors.Is(err, http.ErrServerClosed):
				a.logger.Warn("server shutdown")
			default:
				a.logger.Fatal(err.Error())
			}
		}
		err := a.deps.youtubeServer.Shutdown(ctx)
		if err != nil {
			a.logger.Fatal(err.Error())
		}
	}()

	<-ctx.Done()
}

func (a *app) Close(ctx context.Context) {
	a.deps.Close(ctx, a.logger)
}
