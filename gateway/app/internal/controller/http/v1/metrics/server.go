package metrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VrMolodyakov/vgm/gateway/internal/config"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewServer(cfg config.MetricsServer) *http.Server {
	router := chi.NewRouter()

	router.Get("/metrics", promhttp.Handler().ServeHTTP)

	addr := fmt.Sprintf("%s:%d", cfg.IP, cfg.Port)
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
	}
}
