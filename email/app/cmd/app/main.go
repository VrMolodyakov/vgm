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
	logger.Info(cfg.GRPC)
	logger.Info(cfg.Mail)
	logger.Info(cfg.Nats)
	logger.Info(cfg.Subscriber)
	app := app.NewApp(cfg, logger)
	app.Run(context.Background())

}
