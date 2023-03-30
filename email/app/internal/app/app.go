package app

import (
	"context"
	"fmt"
	"log"

	"github.com/VrMolodyakov/vgm/email/internal/config"
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
