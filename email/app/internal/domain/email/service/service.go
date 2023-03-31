package service

import "github.com/VrMolodyakov/vgm/email/app/internal/domain/email/model"

type EmailClient interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		files []string) error
}

type service struct {
	client EmailClient
}

func NewEmailClient(client EmailClient) *service {
	return &service{
		client: client,
	}
}

func (s *service) Send(email *model.Email) error {
	return s.client.SendEmail(email.Subject, email.Content, email.To, email.Cc, email.Bcc, email.Files)
}
