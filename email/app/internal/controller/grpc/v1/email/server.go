package email

import (
	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/model"
)

type EmailService interface {
	Send(email *model.Email) error
}

type server struct {
	emailService EmailService
	emailPb.UnimplementedEmailServiceServer
}

func NewServer(
	emailService EmailService,
	s emailPb.UnimplementedEmailServiceServer) *server {

	return &server{
		emailService:                    emailService,
		UnimplementedEmailServiceServer: s,
	}
}
