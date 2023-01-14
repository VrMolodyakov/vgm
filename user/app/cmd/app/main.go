package main

import (
	"fmt"
	"log"
	"time"

	"github.com/VrMolodyakov/vgm/user/internal/config"
	"github.com/VrMolodyakov/vgm/user/pkg/logging"
)

func main() {
	logger, err := logging.New("info", "log.txt")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("starting app")
	config := config.GetConfig()
	fmt.Println("config: ", config.Postgres)
	for {
		time.Sleep(10 * time.Minute)
	}
}
