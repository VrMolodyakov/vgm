package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	SuccessRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_email_success_incoming_messages_total",
		Help: "The total number of success incoming email GRPC requests",
	})
	ErrorRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_email_error_incoming_message_total",
		Help: "The total number of error incoming email GRPC requests",
	})
)

func RegisterGrpcMetrics() error {
	if err := prometheus.Register(SuccessRequests); err != nil {
		return err
	}
	if err := prometheus.Register(ErrorRequests); err != nil {
		return err
	}
	return nil
}
