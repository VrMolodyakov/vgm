package main

import (
	"context"

	"github.com/VrMolodyakov/vgm/gateway/internal/app"
	"github.com/VrMolodyakov/vgm/gateway/internal/config"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
)

func main() {
	logging.Init("info", "log.txt")
	logger := logging.GetLogger()
	logger.Info("starting music app")
	cfg := config.GetConfig()

	logger.Sugar().Infow(
		"psql config",
		"database", cfg.Postgres.Database,
		"IP", cfg.Postgres.IP,
		"Password", cfg.Postgres.Password,
		"Pool size", cfg.Postgres.PoolSize,
		"Port", cfg.Postgres.Port,
		"User", cfg.Postgres.User,
	)

	logger.Sugar().Infow(
		"http config",
		"IP", cfg.HTTP.IP,
		"Port", cfg.HTTP.Port,
		"Read TImeout", cfg.HTTP.ReadTimeout,
		"Write Timeout", cfg.HTTP.WriteTimeout,
		"CORS", cfg.HTTP.CORS,
		"User", cfg.Postgres.User,
	)

	logger.Sugar().Infow(
		"redis config",
		"Host", cfg.Redis.Host,
		"Password", cfg.Redis.Password,
		"Port", cfg.Redis.Port,
		"Db number", cfg.Redis.DbNumber,
	)

	// logger.Sugar().Infow(
	// 	"keys config",
	// 	"Access private key", cfg.KeyPairs.AccessPrivate,
	// 	"Access public key", cfg.KeyPairs.AccessPublic,
	// 	"Refresh private key", cfg.KeyPairs.RefreshPrivate,
	// 	"Refresh public key", cfg.KeyPairs.RefreshPublic,
	// 	"Access TTL", cfg.KeyPairs.AccessTtl,
	// 	"Refresh TTL", cfg.KeyPairs.RefreshTtl,
	// )

	ctx := context.Background()
	app := app.NewApp(cfg)
	app.Run(ctx)
}
