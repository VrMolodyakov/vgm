package main

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/app"
	"github.com/VrMolodyakov/vgm/music/app/internal/config"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

func main() {
	logging.Init("info", "log.txt")
	logger := logging.GetLogger()
	logger.Info("starting music app")
	cfg := config.GetConfig()

	logger.Infow(
		"psql config",
		"database", cfg.Postgres.Database,
		"IP", cfg.Postgres.IP,
		"Password", cfg.Postgres.Password,
		"Pool size", cfg.Postgres.PoolSize,
		"Port", cfg.Postgres.Port,
		"User", cfg.Postgres.User,
	)
	ctx := context.Background()
	app := app.NewApp(cfg)
	app.Run(ctx)

}
