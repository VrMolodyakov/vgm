package music

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/config"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/middleware"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewServer(
	cfg config.MusicServer,
	cors middleware.Cors,
	auth *middleware.AuthMiddleware,
	albumService AlbumService,
) *http.Server {

	handler := NewAlbumHandler(albumService)
	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)
	router.Use(cors.CORS)
	router.Use(middleware.DurationMiddleware)
	router.Use(chiMiddleware.Recoverer)

	router.Route("/music", func(r chi.Router) {
		r.Use(auth.Auth)
		r.Post("/album", handler.CreateAlbum)
		r.Post("/person", handler.CreatePerson)
		r.Get("/albums", handler.FindAllAlbums)
		r.Get("/albums/{albumID}", handler.FindFullAlbums)
		r.Get("/dates/{limit}", handler.FindLastUpdateDays)
	})

	addr := fmt.Sprintf("%s:%d", cfg.IP, cfg.Port)
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
	}
}
