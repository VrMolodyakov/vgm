package main

import (
	"context"

	"github.com/VrMolodyakov/vgm/gateway/internal/app"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
)

func main() {
	ctx := context.Background()
	a := app.New()
	logger := logging.GetLogger()

	defer func() {
		logger.Info("shutting down server")
		a.Close(ctx)
		logger.Info("done. exiting")
	}()

	if err := a.ReadConfig(); err != nil {
		logger.Sugar().Error(err, "read config")
		return
	}

	if err := a.InitTracer(); err != nil {
		logger.Sugar().Error(err, "init tracer")
		return
	}

	if err := a.Setup(ctx); err != nil {
		logger.Sugar().Error(err, "setup dependencies")
		return
	}

	a.Start()
}
