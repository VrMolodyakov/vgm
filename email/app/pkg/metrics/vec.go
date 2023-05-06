package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ResponseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_status",
			Help: "Status of HTTP response",
		},
		[]string{"status"},
	)

	AlbumCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Name:      "http_album_counter",
		Help:      "Number of HTTP album responses",
	}, []string{"album", "code"})
)
