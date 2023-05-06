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
)
