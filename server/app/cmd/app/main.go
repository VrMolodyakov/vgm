package main

import (
	"fmt"
	"log"

	"github.com/VrMolodyakov/vgm/internal/config"
	"github.com/VrMolodyakov/vgm/pkg/logging"
)

func main() {
	logger, err := logging.New("info", "log.txt")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("start app")

	config := config.GetConfig()
	fmt.Println(config)
}
