package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/VrMolodyakov/vgm/pkg/logging"
)

func main() {
	logger, err := logging.New("info", "log.txt")
	if err != nil {

	}
	logger.Info("start app")
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("cannot")
	}
	dirname := filepath.Dir(filename)
	logger.Info(dirname)
}
