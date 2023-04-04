package email

import (
	"context"

	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/model"
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
)

type EmailUseCase interface {
	Publush(ctx context.Context, email *model.Email) error
}

type server struct {
	logger      logging.Logger
	emailUseCae EmailUseCase
	emailPb.UnimplementedEmailServiceServer
}

func NewServer(
	emailUseCae EmailUseCase,
	logger logging.Logger,
	s emailPb.UnimplementedEmailServiceServer) *server {

	return &server{
		emailUseCae:                     emailUseCae,
		UnimplementedEmailServiceServer: s,
		logger:                          logger,
	}
}
