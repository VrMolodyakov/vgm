package usecase

import (
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
	"github.com/nats-io/nats.go"
)

type Publisher interface {
	Publish(subject string, data []byte) error
	PublishAsync(subject string, data []byte) (nats.PubAckFuture, error)
}

type EmailClient interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		files []string) error
}

type emailUseCase struct {
	logger      logging.Logger
	publisher   Publisher
	emailClient EmailClient
}

func NewEmailUseCase(
	logger logging.Logger,
	publisher Publisher,
	emailClient EmailClient,
) *emailUseCase {
	return &emailUseCase{
		logger:      logger,
		publisher:   publisher,
		emailClient: emailClient,
	}
}
