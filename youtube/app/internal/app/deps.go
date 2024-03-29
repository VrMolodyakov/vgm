package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	ttlcache "github.com/VrMolodyakov/ttl-cache"
	"github.com/VrMolodyakov/vgm/youtube/internal/config"
	"github.com/VrMolodyakov/vgm/youtube/internal/controller/grpc/v1/client"
	"github.com/VrMolodyakov/vgm/youtube/internal/controller/grpc/v1/client/music"
	"github.com/VrMolodyakov/vgm/youtube/internal/controller/http/v1/middleware"
	youtubeServer "github.com/VrMolodyakov/vgm/youtube/internal/controller/http/v1/youtube"
	musicService "github.com/VrMolodyakov/vgm/youtube/internal/domain/music/service"
	youtubeService "github.com/VrMolodyakov/vgm/youtube/internal/domain/youtube/service"
	"github.com/VrMolodyakov/vgm/youtube/pkg/cache"
	"github.com/VrMolodyakov/vgm/youtube/pkg/logging"
	"github.com/VrMolodyakov/vgm/youtube/pkg/youtube"
)

type Deps struct {
	youtubeServer *http.Server
	cache         *ttlcache.Cache[string, string]
}

func (d *Deps) Setup(ctx context.Context, cfg *config.Config, logger logging.Logger) error {
	musicAddress := fmt.Sprintf("%s:%d", cfg.MusicGRPC.HostName, cfg.MusicGRPC.Port)
	musicCerts := d.loadYoutubeClientCert(cfg.YoutubeClientCert)
	grpcMusicClient := music.NewMusicClient(musicAddress, logger)
	grpcMusicClient.StartWithTLS(musicCerts)
	music := musicService.NewMusicService(grpcMusicClient, logger)

	client, err := youtube.NewYoutubeClient(ctx, cfg.Youtube.ApiKey)
	if err != nil {
		log.Fatal(err)
	}
	youtube := youtubeService.NewYoutubeService(client)
	cors := d.setupCORS(cfg.CORS)
	d.cache = cache.New[string, string](cfg.Cache.CleanInterval)
	d.cache.Clean()
	d.youtubeServer = youtubeServer.NewServer(
		cfg.YoutubeServer,
		logger,
		cors,
		music,
		youtube,
		d.cache,
		cfg.Cache.ExpireAt,
	)

	return nil
}

func (d *Deps) Close(ctx context.Context, logger logging.Logger) {
	if d.youtubeServer != nil {
		if err := d.youtubeServer.Shutdown(ctx); err != nil {
			logger.Error(err, "shutdown user server")
		}
	}
	d.cache.Close()
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
