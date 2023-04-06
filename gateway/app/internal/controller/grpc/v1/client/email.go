package client

import (
	"context"
	"log"

	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/email/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type emailClient struct {
	target string
	client emailPb.EmailServiceClient
}

func NewEmailClient(target string) *emailClient {
	if target == "" {
		log.Fatalln("Error in Access to GRPC URL in music client")
	}
	return &emailClient{
		target: target,
	}
}

func (e *emailClient) Start() {

	conn, err := grpc.Dial(e.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	e.client = emailPb.NewEmailServiceClient(conn)
}

func (e *emailClient) Send(ctx context.Context, email model.Email) error {
	logger := logging.LoggerFromContext(ctx)
	req := emailPb.CreateEmailRequest{
		Subject: email.Subject,
		Content: email.Content,
		Bcc:     email.Bcc,
		Cc:      email.Cc,
		To:      email.To,
		Files:   email.Files,
	}
	_, err := e.client.CreateEmail(ctx, &req)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
