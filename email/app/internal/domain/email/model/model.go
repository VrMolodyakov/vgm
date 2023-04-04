package model

import (
	"time"

	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	"github.com/VrMolodyakov/vgm/email/app/pkg/errors"
)

type Email struct {
	Subject string
	Content string
	To      []string
	Cc      []string
	Bcc     []string
	Files   []string
}

type EmailErrorMsg struct {
	Subject string
	Reply   string
	Data    []byte
	Error   error
	Time    time.Time
}

func ModelFromPB(req *emailPb.CreateEmailRequest) *Email {
	return &Email{
		Subject: req.GetSubject(),
		Content: req.GetContent(),
		To:      req.GetTo(),
		Cc:      req.GetCc(),
		Bcc:     req.GetBcc(),
		Files:   req.GetFiles(),
	}
}

func (e *Email) Validate() error {
	if len(e.Subject) == 0 {
		return errors.New("Validation error: Subject")
	}
	if len(e.Content) == 0 {
		return errors.New("Validation error: Content")
	}
	if len(e.To) == 0 {
		return errors.New("Validation error: To")
	}
	return nil
}
