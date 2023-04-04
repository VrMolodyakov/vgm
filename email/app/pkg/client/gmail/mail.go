package gmail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type gmailClient struct {
	smtpAuthAddress   string
	smtpServerAddress string
	name              string
	fromAddress       string
	fromPassword      string
}

func NewMailClient(
	smtpAuthAddress string,
	smtpServerAddress string,
	name string,
	fromAddress string,
	fromPassword string) *gmailClient {
	return &gmailClient{
		smtpAuthAddress:   smtpAuthAddress,
		smtpServerAddress: smtpServerAddress,
		name:              name,
		fromAddress:       fromAddress,
		fromPassword:      fromPassword,
	}
}

func (g *gmailClient) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string) error {

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", g.name, g.fromAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", g.fromAddress, g.fromPassword, g.smtpAuthAddress)
	return e.Send(g.smtpServerAddress, smtpAuth)
}
