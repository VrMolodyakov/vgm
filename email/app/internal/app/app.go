package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/VrMolodyakov/vgm/email/app/internal/config"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

const (
	subjectName string = "email"
)

type app struct {
	cfg        *config.Config
	grpcServer *grpc.Server
}

func NewApp(cfg *config.Config) *app {
	return &app{cfg: cfg}
}

func (a *app) Run(ctx context.Context) {
	a.startGrpc(ctx)
}

func (a *app) startGrpc(ctx context.Context) {
	logger := logging.LoggerFromContext(ctx)
	logger.Infow("grpc cfg ", "gprc ip : ", a.cfg.GRPC.IP, "gprc port :", a.cfg.GRPC.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		logger.Error(err.Error())
	}
}

func (a *app) createStreamContext() nats.JetStreamContext {
	address := fmt.Sprintf("nats://%s:%d", a.cfg.Nats.Host, a.cfg.Nats.Port)
	n, err := nats.Connect(address)
	if err != nil {
		log.Fatal(err)
	}
	streamContext, err := n.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}
	s, err := streamContext.StreamInfo(subjectName)
	if s == nil {
		// Create new stream
		_, err := streamContext.AddStream(&nats.StreamConfig{
			Name:     subjectName,
			Subjects: []string{"email.*"},
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	return streamContext
}
