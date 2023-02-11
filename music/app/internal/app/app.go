package app

import (
	"context"
	"fmt"
	"net"
	"time"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	creditPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/credit/v1"
	infoPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/info/v1"
	personPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/person/v1"

	"github.com/VrMolodyakov/vgm/music/app/internal/config"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/v1/album"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/v1/credit"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/v1/info"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/v1/person"

	albumDAO "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/dao"
	AlbumPolicy "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/policy"
	AlbumService "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/service"
	creditDAO "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/dao"
	creditPolicy "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/policy"
	creditService "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/service"
	infoDAO "github.com/VrMolodyakov/vgm/music/app/internal/domain/info/dao"
	infoPolicy "github.com/VrMolodyakov/vgm/music/app/internal/domain/info/policy"
	infoService "github.com/VrMolodyakov/vgm/music/app/internal/domain/info/service"
	personDAO "github.com/VrMolodyakov/vgm/music/app/internal/domain/person/dao"
	personPolicy "github.com/VrMolodyakov/vgm/music/app/internal/domain/person/policy"
	personService "github.com/VrMolodyakov/vgm/music/app/internal/domain/person/service"

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

	albumDAO := albumDAO.NewAlbumStorage(pgClient)
	albumService := AlbumService.NewAlbumService(albumDAO)
	albumPolicy := AlbumPolicy.NewAlbumPolicy(albumService)

	infoDAO := infoDAO.NewInfoStorage(pgClient)
	infoService := infoService.NewInfoService(infoDAO)
	infoPolicy := infoPolicy.NewInfoPolicy(infoService)
	infoServer := info.NewServer(infoPolicy, infoPb.UnimplementedInfoServiceServer{})

	albumServer := album.NewServer(albumPolicy, albumPb.UnimplementedAlbumServiceServer{})

	creditDAO := creditDAO.NewCreditStorage(pgClient)
	creditService := creditService.NewCreditService(creditDAO)
	creditPolicy := creditPolicy.NewCreditPolicy(creditService)

	creditServer := credit.NewServer(creditPolicy, creditPb.UnimplementedCreditServiceServer{})

	personDAO := personDAO.NewPersonStorage(pgClient)
	personService := personService.NewPersonService(personDAO)
	personPolicy := personPolicy.NewPersonPolicy(personService)

	personServer := person.NewServer(personPolicy, personPb.UnimplementedPersonServiceServer{})

	a.grpcServer = grpc.NewServer(serverOptions...)

	albumPb.RegisterAlbumServiceServer(a.grpcServer, albumServer)
	infoPb.RegisterInfoServiceServer(a.grpcServer, infoServer)
	creditPb.RegisterCreditServiceServer(a.grpcServer, creditServer)
	personPb.RegisterPersonServiceServer(a.grpcServer, personServer)

	reflection.Register(a.grpcServer)
	logger.Info("start grpc serve")
	a.grpcServer.Serve(listener)
	logger.Info("end of gprc")

}
