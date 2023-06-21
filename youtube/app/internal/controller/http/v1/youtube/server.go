package youtube

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VrMolodyakov/vgm/youtube/internal/config"
	"github.com/VrMolodyakov/vgm/youtube/internal/controller/http/v1/middleware"
	"github.com/VrMolodyakov/vgm/youtube/pkg/logging"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewServer(
	cfg config.YoutubeServer,
	logger logging.Logger,
	cors middleware.Cors,
	music MusicService,
	youtube YoutubeService,
) *http.Server {

	handler := NewYoutubeHandler(logger, youtube, music)
	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)
	router.Use(cors.CORS)
	router.Use(chiMiddleware.Recoverer)

	router.Route("/youtube", func(r chi.Router) {
		r.Get("/playlist/{limit}", handler.CreatePlaylist)
	})

	addr := fmt.Sprintf("%s:%d", cfg.IP, cfg.Port)
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
	}
}
