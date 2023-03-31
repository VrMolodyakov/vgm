package model

import (
	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
)

type Email struct {
	Subject string
	Content string
	To      []string
	Cc      []string
	Bcc     []string
	Files   []string
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
