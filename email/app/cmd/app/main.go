package main

import (
	"fmt"

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
}
