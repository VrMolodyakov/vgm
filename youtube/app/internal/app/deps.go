package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/VrMolodyakov/vgm/youtube/internal/config"
	"github.com/VrMolodyakov/vgm/youtube/internal/controller/grpc/v1/client"
	"github.com/VrMolodyakov/vgm/youtube/internal/controller/grpc/v1/client/music"
	"github.com/VrMolodyakov/vgm/youtube/internal/controller/http/v1/middleware"
	youtubeServer "github.com/VrMolodyakov/vgm/youtube/internal/controller/http/v1/youtube"
	musicService "github.com/VrMolodyakov/vgm/youtube/internal/domain/music/service"
	youtubeService "github.com/VrMolodyakov/vgm/youtube/internal/domain/youtube/service"
	"github.com/VrMolodyakov/vgm/youtube/pkg/logging"
	"github.com/VrMolodyakov/vgm/youtube/pkg/youtube"
)

type Deps struct {
	youtubeServer *http.Server
}

func (d *Deps) Setup(ctx context.Context, cfg *config.Config, logger logging.Logger) error {
	musicAddress := fmt.Sprintf("%s:%d", cfg.MusicGRPC.HostName, cfg.MusicGRPC.Port)
	grpcMusicClient := music.NewMusicClient(musicAddress)
	music := musicService.NewMusicService(grpcMusicClient, logger)

	client, err := youtube.NewYoutubeClient(ctx, "")
	if err != nil {
		log.Fatal(err)
	}
	youtube := youtubeService.NewYoutubeService(client)
	cors := d.setupCORS(cfg.CORS)
	d.youtubeServer = youtubeServer.NewServer(cfg.YoutubeServer, logger, cors, music, youtube)

	return nil
}

func (d *Deps) Close(ctx context.Context, logger logging.Logger) {
	if d.youtubeServer != nil {
		if err := d.youtubeServer.Shutdown(ctx); err != nil {
			logger.Error(err, "shutdown user server")
		}
	}
}

func (d *Deps) loadYoutubeClientCert(cfg config.YoutubeClientCert) client.ClientCerts {
	return client.NewClientCerts(
		cfg.EnableTLS,
		cfg.ClientCertFile,
		cfg.ClientKeyFile,
		cfg.ClientCACertFile,
	)
}

func (d *Deps) setupCORS(cfg config.CORS) middleware.Cors {
	origins := strings.Join(cfg.AllowedOrigins[:], ", ")
	headers := strings.Join(cfg.AllowedHeaders[:], ", ")
	methods := strings.Join(cfg.AllowedMethods[:], ", ")

	return middleware.NewCors(origins, headers, methods)

}
