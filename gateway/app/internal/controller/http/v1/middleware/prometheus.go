package middleware

import (
	"net/http"

	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

func DurationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		timer := prometheus.NewTimer(metrics.HttpDuration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}
