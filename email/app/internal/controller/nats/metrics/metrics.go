package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TotalSubscribeMessages = promauto.NewCounter(prometheus.CounterOpts{
		Name: "nats_email_incoming_messages_total",
		Help: "The total number of incoming email NATS messages",
	})
	SuccessSubscribeMessages = promauto.NewCounter(prometheus.CounterOpts{
		Name: "nats_email_success_incoming_messages_total",
		Help: "The total number of success email NATS messages",
	})
	ErrorSubscribeMessages = promauto.NewCounter(prometheus.CounterOpts{
		Name: "nats_email_error_incoming_messages_total",
		Help: "The total number of error email NATS messages",
	})
)

func RegisterNatsMetrics() error {
	if err := prometheus.Register(TotalSubscribeMessages); err != nil {
		return err
	}
	if err := prometheus.Register(SuccessSubscribeMessages); err != nil {
		return err
	}
	if err := prometheus.Register(ErrorSubscribeMessages); err != nil {
		return err
	}
	return nil
}
