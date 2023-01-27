package app

import (
	"context"
	"fmt"
	"net"
	"time"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"

	"github.com/VrMolodyakov/vgm/music/app/internal/config"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/v1/album"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/policy"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/service"
	"github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type app struct {
	cfg        *config.Config
	grpcServer *grpc.Server
}

func NewApp(cfg *config.Config) *app {
	return &app{cfg: cfg}
}

func (a *app) Run(ctx context.Context) {
	a.startGrpc(ctx)
}

func (a *app) startGrpc(ctx context.Context) {
	logger := logging.LoggerFromContext(ctx)
	logger.Infow("grpc cfg ", "gprc ip : ", a.cfg.GRPC.IP, "gprc port :", a.cfg.GRPC.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		logger.Error(err.Error())
	}
	serverOptions := []grpc.ServerOption{}

	pgConfig := postgresql.NewPgConfig(
		a.cfg.Postgres.User,
		a.cfg.Postgres.Password,
		a.cfg.Postgres.IP,
		a.cfg.Postgres.Port,
		a.cfg.Postgres.Database,
		a.cfg.Postgres.PoolSize)

	pgClient, err := postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	storage := dao.NewProductStorage(pgClient)
	service := service.NewAlbumService(storage)
	policy := policy.NewAlbumPolicy(service)
	albumServer := album.NewServer(policy, albumPb.UnimplementedAlbumServiceServer{})

	a.grpcServer = grpc.NewServer(serverOptions...)

	albumPb.RegisterAlbumServiceServer(a.grpcServer, albumServer)

	reflection.Register(a.grpcServer)
	logger.Info("start grpc serve")
	a.grpcServer.Serve(listener)
	logger.Info("end of gprc")

}
