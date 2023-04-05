package app

import (
	"context"
	"fmt"

	"net"

	"github.com/VrMolodyakov/vgm/email/app/internal/config"
	"github.com/VrMolodyakov/vgm/email/app/internal/controller/grpc/v1/email"
	"github.com/VrMolodyakov/vgm/email/app/internal/controller/grpc/v1/interceptor"

	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	jet "github.com/VrMolodyakov/vgm/email/app/internal/controller/nats"
	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/usecase"
	"github.com/VrMolodyakov/vgm/email/app/pkg/client/gmail"
	"github.com/VrMolodyakov/vgm/email/app/pkg/client/nats"
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	subjectName string = "email"
)

type app struct {
	cfg        *config.Config
	logger     logging.Logger
	grpcServer *grpc.Server
}

func NewApp(cfg *config.Config, logger logging.Logger) *app {
	return &app{cfg: cfg, logger: logger}
}

func (a *app) Run(ctx context.Context) {
	streamCtx := nats.NewStreamContext(
		a.cfg.Nats.Host,
		a.cfg.Nats.Port,
		a.cfg.Subscriber.MainSubjectName,
		a.cfg.Subscriber.MainSubjects,
	)
	info, err := streamCtx.StreamInfo("email")
	if err != nil {
		a.logger.Fatal(err)
	}
	a.logger.Info("cluster", info.Cluster)
	a.logger.Info("config", info.Config)
	a.logger.Info("config", info.Sources)
	pub := jet.NewPublisher(streamCtx)
	emailClient := gmail.NewMailClient(
		a.cfg.Mail.SmtpAuthAddress,
		a.cfg.Mail.SmtpServerAddress,
		a.cfg.Mail.Name,
		a.cfg.Mail.FromAddress,
		a.cfg.Mail.FromPassword,
	)
	emailUseCase := usecase.NewEmailUseCase(a.logger, pub, a.cfg.Subscriber.SendEmailSubject, emailClient)
	go func() {
		subCfg := jet.NewSubscriberCfg(
			a.cfg.Subscriber.DurableName,
			a.cfg.Subscriber.DeadMessageSubject,
			a.cfg.Subscriber.SendEmailSubject,
			a.cfg.Subscriber.EmailGroupName,
			a.cfg.Subscriber.AckWait,
			a.cfg.Subscriber.Workers,
			a.cfg.Subscriber.MaxInflight,
			a.cfg.Subscriber.MaxDeliver,
		)
		sub := jet.NewSubscriber(streamCtx, emailUseCase, subCfg, a.logger)
		sub.Run(ctx)
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	fmt.Printf("print ip = %s port = %d", a.cfg.GRPC.IP, a.cfg.GRPC.Port)
	a.logger.Info("grpc listener :=", zap.String("ip", a.cfg.GRPC.IP), zap.Int("port", a.cfg.GRPC.Port))
	if err != nil {
		a.logger.Error(err.Error())
	}

	emailServer := email.NewServer(emailUseCase, a.logger, emailPb.UnimplementedEmailServiceServer{})

	a.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.NewLoggerInterceptor(a.logger),
		),
	)
	emailPb.RegisterEmailServiceServer(a.grpcServer, emailServer)
	reflection.Register(a.grpcServer)
	a.logger.Info("start grpc serve")
	a.grpcServer.Serve(listener)
	a.logger.Info("end of email service")
}
