package main

import (
	"fmt"

	"github.com/VrMolodyakov/vgm/music/app/internal/config"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

func main() {
	logger, err := logging.New("info", "log.txt")
	if err != nil {
		fmt.Errorf("cannot start app due to %v", err)
	}
	logger.Info("music service start")
	cfg := config.GetConfig()
	logger.Sugar().Info(cfg)
	logger.Info("music service end")

}
