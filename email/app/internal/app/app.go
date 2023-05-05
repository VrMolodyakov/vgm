package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"net"

	"github.com/VrMolodyakov/vgm/email/app/internal/config"

	"github.com/VrMolodyakov/vgm/email/app/pkg/jaeger"
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
	"google.golang.org/grpc/reflection"
)

//TODO:put in yaml
const (
	enableTLS               = true
	serverCertFile   string = "cert/email-server-cert.pem"
	serverKeyFile    string = "cert/email-server-key.pem"
	clientCACertFile string = "cert/ca-cert.pem"
	serviceName      string = "email-server"
)

type app struct {
	cfg    *config.Config
	logger logging.Logger
	deps   Deps
}

func New() *app {
	return &app{}
}

func (a *app) Setup() error {
	return a.deps.Setup(a.cfg, a.logger)
}

func (a *app) InitLogger() {
	a.logger = logging.NewLogger(a.cfg.Logger)
	a.logger.InitLogger()
}

func (a *app) InitTracer() error {
	err := jaeger.SetGlobalTracer(a.cfg.Jaeger.ServiceName, a.cfg.Jaeger.Address, a.cfg.Jaeger.Port)
	if err != nil {
		return err
	}
	return nil
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

func (a *app) Start(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, os.Kill, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	a.logger.Infow("grpc cfg ", "gprc ip : ", a.cfg.GRPC.IP, "gprc port :", a.cfg.GRPC.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		reflection.Register(a.deps.server)
		a.logger.Info("music grpc server started...")
		a.deps.server.Serve(listener)
		a.logger.Info("end of music gprc server")

	}()

	<-ctx.Done()
}

func (a *app) StartSubscriber(ctx context.Context) {
	a.deps.subscriber.Run(ctx)
}

func (a *app) PrintConfig() {
	a.logger.Infow(
		"grpc cfg: ",
		"ip", a.cfg.GRPC.IP,
		"port", a.cfg.GRPC.Port,
	)

	a.logger.Infow(
		"mail cfg: ",
		"from", a.cfg.Mail.FromAddress,
		"password", a.cfg.Mail.FromPassword,
		"name", a.cfg.Mail.Name,
		"smtp address", a.cfg.Mail.SmtpAuthAddress,
		"smtp server address", a.cfg.Mail.SmtpServerAddress,
	)

	a.logger.Infow(
		"nats cfg: ",
		"host", a.cfg.Nats.Host,
		"port", a.cfg.Nats.Port,
	)

	a.logger.Infow(
		"subscriber cfg: ",
		"ackWait", a.cfg.Subscriber.AckWait,
		"deadMessageSubject", a.cfg.Subscriber.DeadMessageSubject,
		"durableName", a.cfg.Subscriber.DurableName,
		"emailGroupName", a.cfg.Subscriber.EmailGroupName,
		"mainSubjectName", a.cfg.Subscriber.MainSubjectName,
		"mainSubject", a.cfg.Subscriber.MainSubjects,
		"maxDeliver", a.cfg.Subscriber.MaxDeliver,
		"maxInflight", a.cfg.Subscriber.MaxInflight,
		"sendEmailSubject", a.cfg.Subscriber.SendEmailSubject,
		"workers", a.cfg.Subscriber.Workers,
	)
}
