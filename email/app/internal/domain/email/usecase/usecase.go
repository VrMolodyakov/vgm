package usecase

import (
	"context"
	"encoding/json"

	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/model"
	"github.com/VrMolodyakov/vgm/email/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("email-usecase")
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

type EmailUseCase struct {
	logger      logging.Logger
	sendSubject string
	publisher   Publisher
	emailClient EmailClient
}

func NewEmailUseCase(
	logger logging.Logger,
	publisher Publisher,
	subject string,
	emailClient EmailClient,
) *EmailUseCase {
	return &EmailUseCase{
		logger:      logger,
		sendSubject: subject,
		publisher:   publisher,
		emailClient: emailClient,
	}
}

func (e *EmailUseCase) Publush(ctx context.Context, email *model.Email) error {
	_, span := tracer.Start(ctx, "emailUseCase.Publush")
	defer span.End()

	mailBytes, err := json.Marshal(email)
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}
	return e.publisher.Publish(e.sendSubject, mailBytes)
}

func (e *EmailUseCase) Send(ctx context.Context, email *model.Email) error {
	_, span := tracer.Start(ctx, "emailUseCase.Send")
	defer span.End()

	if err := e.emailClient.SendEmail(
		email.Subject,
		email.Content,
		email.To,
		email.Cc,
		email.Bcc,
		email.Files,
	); err != nil {
		return errors.Wrap(err, "SendEmail")
	}

	return nil
}
