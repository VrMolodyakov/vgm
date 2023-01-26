package app

import (
	"context"
	"net/http"

	"github.com/VrMolodyakov/vgm/music/app/internal/config"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
)

type App struct {
	cfg *config.Config

	router     *httprouter.Router
	httpServer *http.Server
	grpcServer *grpc.Server

	pgClient *pgxpool.Pool
}

func (a *App) Run(ctx context.Context) {

}

func (a *App) startGrpc(ctx context.Context) {
	logger := logging.LoggerFromContext(ctx)
	logger.With("gprc ip : ", a.cfg.GRPC.IP, "gprc port : ", a.cfg.GRPC.Port)

}
