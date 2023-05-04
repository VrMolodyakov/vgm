package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/VrMolodyakov/vgm/music/app/internal/config"

	"github.com/VrMolodyakov/vgm/music/app/pkg/jaeger"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"google.golang.org/grpc/reflection"
)

//TODO:put in yaml
const (
	enableTLS               = true
	serverCertFile   string = "cert/server-cert.pem"
	serverKeyFile    string = "cert/server-key.pem"
	clientCACertFile string = "cert/ca-cert.pem"
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

func (a *app) Setup(ctx context.Context) error {
	return a.deps.Setup(ctx, a.cfg)
}

func (a *app) Close() {
	a.deps.Close()
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

	logger := logging.GetLogger()
	logger.Infow("grpc cfg ", "gprc ip : ", a.cfg.GRPC.IP, "gprc port :", a.cfg.GRPC.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		reflection.Register(a.deps.server)
		logger.Info("music grpc server started...")
		a.deps.server.Serve(listener)
		logger.Info("end of music gprc server")

	}()

	<-ctx.Done()
}
