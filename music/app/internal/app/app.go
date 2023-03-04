package app

import (
	"context"
	"fmt"
	"net"
	"time"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	personPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/person/v1"

	"github.com/VrMolodyakov/vgm/music/app/internal/config"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/v1/album"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/v1/person"

	AlbumPolicy "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/policy"
	albumRepository "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/repository"
	AlbumService "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/service"
	creditRepository "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/repository"
	creditService "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/service"

	personPolicy "github.com/VrMolodyakov/vgm/music/app/internal/domain/person/policy"
	personRepository "github.com/VrMolodyakov/vgm/music/app/internal/domain/person/repository"
	personService "github.com/VrMolodyakov/vgm/music/app/internal/domain/person/service"
	tracklistRepository "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/repository"
	tracklistService "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/service"

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
		a.cfg.Postgres.PoolSize,
	)

	pgClient, err := postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}

	creditRepository := creditRepository.NewCreditRepo(pgClient)
	creditService := creditService.NewCreditService(creditRepository)

	personRepository := personRepository.NewPersonStorage(pgClient)
	personService := personService.NewPersonService(personRepository)
	personPolicy := personPolicy.NewPersonPolicy(personService)

	personServer := person.NewServer(personPolicy, personPb.UnimplementedPersonServiceServer{})

	trackRepository := tracklistRepository.NewTracklistRepo(pgClient)
	trackService := tracklistService.NewTrackService(trackRepository)

	albumRepository := albumRepository.NewAlbumRepository(pgClient)
	albumService := AlbumService.NewAlbumService(albumRepository)
	albumPolicy := AlbumPolicy.NewAlbumPolicy(albumService, trackService, creditService)

	albumServer := album.NewServer(albumPolicy, albumPb.UnimplementedAlbumServiceServer{})

	a.grpcServer = grpc.NewServer(serverOptions...)

	albumPb.RegisterAlbumServiceServer(a.grpcServer, albumServer)
	personPb.RegisterPersonServiceServer(a.grpcServer, personServer)

	reflection.Register(a.grpcServer)
	logger.Info("start grpc serve")
	a.grpcServer.Serve(listener)
	logger.Info("end of gprc")

}
