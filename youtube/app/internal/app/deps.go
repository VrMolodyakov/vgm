package app

import (
	"github.com/VrMolodyakov/vgm/youtube/internal/config"
	"github.com/VrMolodyakov/vgm/youtube/internal/controller/grpc/v1/client"
)

type Deps struct {
}

func (d *Deps) loadYoutubeClientCert(cfg config.YoutubeClientCert) client.ClientCerts {
	return client.NewClientCerts(
		cfg.EnableTLS,
		cfg.ClientCertFile,
		cfg.ClientKeyFile,
		cfg.ClientCACertFile,
	)
}
