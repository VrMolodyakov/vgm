package app

import (
	"context"
	"fmt"
	"net"

	"github.com/VrMolodyakov/vgm/email/app/internal/config"
	"github.com/VrMolodyakov/vgm/email/app/internal/controller/grpc/v1/interceptor"
	jet "github.com/VrMolodyakov/vgm/email/app/internal/controller/nats"
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
	streamCtx := nats.NewStreamContext(a.cfg.Nats.Host, a.cfg.Nats.Port, a.cfg.Subscriber.MainSubject, a.cfg.Subscriber.MainSubjects)
	jet.NewPublisher(streamCtx)
	jet.NewSubscriberCfg(
		a.cfg.Subscriber.DurableName,
		a.cfg.Subscriber.DeadMessageSubject,
		a.cfg.Subscriber.SendEmailSubject,
		a.cfg.Subscriber.EmailGroupName,
		a.cfg.Subscriber.AckWait,
		a.cfg.Subscriber.Workers,
		a.cfg.Subscriber.MaxInflight,
		a.cfg.Subscriber.MaxDeliver,
	)
	fmt.Println("end of email service")
	// a.startGrpc(ctx)
}

func (a *app) startGrpc(ctx context.Context) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	a.logger.Info("grpc listener :=", zap.String("ip", a.cfg.GRPC.IP), zap.Int("port", a.cfg.GRPC.Port))
	if err != nil {
		a.logger.Error(err.Error())
	}
	a.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.NewLoggerInterceptor(a.logger),
		),
	)
	reflection.Register(a.grpcServer)
	a.logger.Info("start grpc serve")
	a.grpcServer.Serve(listener)
	a.logger.Info("end of gprc")
}

// func (a *app) createStreamContext() nats.JetStreamContext {
// 	address := fmt.Sprintf("nats://%s:%d", a.cfg.Nats.Host, a.cfg.Nats.Port)
// 	a.logger.Info("address := ", address)
// 	n, err := nats.Connect(address)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	streamContext, err := n.JetStream(nats.PublishAsyncMaxPending(256))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = streamContext.AddStream(&nats.StreamConfig{
// 		Name:     subjectName,
// 		Subjects: []string{"email.*"},
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return streamContext
// }
