package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	UserCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_create_user_success_total",
		Help: "The total number of success incoming create user HTTP requests",
	})
	ErrorCreateUserRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_email_error_incoming_message_total",
		Help: "The total number of error incoming create user HTTP requests",
	})
	HttpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
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

func RegisterMetrics() error {
	if err := prometheus.Register(ErrorCreateUserRequests); err != nil {
		return err
	}
	if err := prometheus.Register(UserCounter); err != nil {
		return err
	}
	if err := prometheus.Register(HttpDuration); err != nil {
		return err
	}
	if err := prometheus.Register(AlbumCounter); err != nil {
		return err
	}
	return nil
}
