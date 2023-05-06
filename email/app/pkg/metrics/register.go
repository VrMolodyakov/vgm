package metrics

import "github.com/prometheus/client_golang/prometheus"

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
