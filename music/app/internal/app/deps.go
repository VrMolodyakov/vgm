package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/config"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/v1/album"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/v1/interceptor"
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
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Deps struct {
	server       *grpc.Server
	postgresPool *pgxpool.Pool
}

func (d *Deps) Setup(ctx context.Context, cfg *config.Config) error {
	logger := logging.LoggerFromContext(ctx)
	logger.Info("Setup...")
	var err error

	pgConfig := postgresql.NewPgConfig(
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.IP,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
		cfg.Postgres.PoolSize,
	)

	d.postgresPool, err = postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		return err
	}

	creditRepository := creditRepository.NewCreditRepo(d.postgresPool)
	creditService := creditService.NewCreditService(creditRepository)

	personRepository := personRepository.NewPersonStorage(d.postgresPool)
	personService := personService.NewPersonService(personRepository)
	personPolicy := personPolicy.NewPersonPolicy(personService)

	trackRepository := tracklistRepository.NewTracklistRepo(d.postgresPool)
	trackService := tracklistService.NewTrackService(trackRepository)

	albumRepository := albumRepository.NewAlbumRepository(d.postgresPool)
	albumService := AlbumService.NewAlbumService(albumRepository)
	albumPolicy := AlbumPolicy.NewAlbumPolicy(albumService, trackService, creditService)

	albumServer := album.NewServer(albumPolicy, personPolicy, albumPb.UnimplementedMusicServiceServer{})

	serverOptions := []grpc.ServerOption{}
	if enableTLS {
		tlsCredentials, err := d.loadTLSCredentials()
		if err != nil {
			logger.Sugar().Fatalf("cannot load TLS credentials: %s", err.Error())
		}

		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}

	serverOptions = append(serverOptions, grpc.ChainUnaryInterceptor(
		interceptor.NewLoggerInterceptor(logging.GetLogger().Logger),
	))

	d.server = grpc.NewServer(serverOptions...)

	albumPb.RegisterMusicServiceServer(d.server, albumServer)

	return nil
}

func (d *Deps) Close() {
	if d.server != nil {
		d.server.Stop()
	}

	if d.postgresPool != nil {
		d.postgresPool.Close()
	}
}

func (d *Deps) loadTLSCredentials() (credentials.TransportCredentials, error) {
	dockerPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	containerConfigPath := filepath.Dir(filepath.Dir(dockerPath))
	path := containerConfigPath + clientCACertFile
	pemClientCA, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	serverCert, err := tls.LoadX509KeyPair(containerConfigPath+serverCertFile, containerConfigPath+serverKeyFile)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}
