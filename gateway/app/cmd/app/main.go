package main

import (
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
)

func main() {
	logging.Init("info", "log.txt")
	logger := logging.GetLogger()
	logger.Info("user service star")
	logger.Info("user service end")
}
