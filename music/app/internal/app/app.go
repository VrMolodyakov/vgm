package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
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
	"github.com/VrMolodyakov/vgm/music/app/pkg/jaeger"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

//TODO:put in yaml
const (
	enableTLS               = true
	serverCertFile   string = "cert/server-cert.pem"
	serverKeyFile    string = "cert/server-key.pem"
	clientCACertFile string = "cert/ca-cert.pem"
	serviceName      string = "music-service"
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

	trackRepository := tracklistRepository.NewTracklistRepo(pgClient)
	trackService := tracklistService.NewTrackService(trackRepository)

	albumRepository := albumRepository.NewAlbumRepository(pgClient)
	albumService := AlbumService.NewAlbumService(albumRepository)
	albumPolicy := AlbumPolicy.NewAlbumPolicy(albumService, trackService, creditService)

	albumServer := album.NewServer(albumPolicy, personPolicy, albumPb.UnimplementedMusicServiceServer{})

	err = jaeger.SetGlobalTracer(serviceName, a.cfg.Jaeger.Address, a.cfg.Jaeger.Port)
	if err != nil {
		logger.Fatal(err.Error())
	}

	serverOptions := []grpc.ServerOption{}
	if enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			logger.Sugar().Fatalf("cannot load TLS credentials: %s", err.Error())
		}

		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}

	serverOptions = append(serverOptions, grpc.ChainUnaryInterceptor(
		interceptor.NewLoggerInterceptor(logging.GetLogger().Logger),
	))

	err = jaeger.SetGlobalTracer(serviceName, a.cfg.Jaeger.Address, a.cfg.Jaeger.Port)
	if err != nil {
		logger.Fatal(err.Error())
	}

	a.grpcServer = grpc.NewServer(serverOptions...)

	albumPb.RegisterMusicServiceServer(a.grpcServer, albumServer)

	reflection.Register(a.grpcServer)
	logger.Info("start grpc serve")
	a.grpcServer.Serve(listener)
	logger.Info("end of gprc")

}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
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
