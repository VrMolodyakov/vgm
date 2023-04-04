package main

import (
	"context"
	"fmt"

	"github.com/VrMolodyakov/vgm/email/app/internal/app"
	"github.com/VrMolodyakov/vgm/email/app/internal/config"
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println(cfg.Logger)
	logger := logging.NewLogger(cfg)
	logger.InitLogger()
	logger.Info("start email service")

	logger.Infow(
		"grpc cfg: ",
		"ip", cfg.GRPC.IP,
		"port", cfg.GRPC.Port,
	)

	logger.Infow(
		"mail cfg: ",
		"from", cfg.Mail.FromAddress,
		"password", cfg.Mail.FromPassword,
		"name", cfg.Mail.Name,
		"smtp address", cfg.Mail.SmtpAuthAddress,
		"smtp server address", cfg.Mail.SmtpServerAddress,
	)

	logger.Infow(
		"nats cfg: ",
		"host", cfg.Nats.Host,
		"port", cfg.Nats.Port,
	)

	logger.Infow(
		"subscriber cfg: ",
		"ackWait", cfg.Subscriber.AckWait,
		"deadMessageSubject", cfg.Subscriber.DeadMessageSubject,
		"durableName", cfg.Subscriber.DurableName,
		"emailGroupName", cfg.Subscriber.EmailGroupName,
		"mainSubjectName", cfg.Subscriber.MainSubjectName,
		"mainSubject", cfg.Subscriber.MainSubjects,
		"maxDeliver", cfg.Subscriber.MaxDeliver,
		"maxInflight", cfg.Subscriber.MaxInflight,
		"sendEmailSubject", cfg.Subscriber.SendEmailSubject,
		"workers", cfg.Subscriber.Workers,
	)

	app := app.NewApp(cfg, logger)
	app.Run(context.Background())

}
