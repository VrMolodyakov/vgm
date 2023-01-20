package main

import (
	"log"

	"github.com/VrMolodyakov/vgm/user/pkg/logging"
)

func main() {
	logger, err := logging.New("info", "log.txt")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("user service star")
	logger.Info("user service end")
}
