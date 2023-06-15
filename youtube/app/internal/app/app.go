package app

import (
	"github.com/VrMolodyakov/vgm/youtube/internal/config"
	"github.com/VrMolodyakov/vgm/youtube/pkg/jaeger"
	"github.com/VrMolodyakov/vgm/youtube/pkg/logging"
)

const (
	enableTLS               = true
	serverCertFile   string = "cert/yt-server-cert.pem"
	serverKeyFile    string = "cert/yt-server-key.pem"
	clientCACertFile string = "cert/ca-cert.pem"
)

type app struct {
	cfg    *config.Config
	logger logging.Logger
	deps   Deps
}

func New() *app {
	return &app{}
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
