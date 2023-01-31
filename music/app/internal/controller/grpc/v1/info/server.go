package info

import (
	"context"

	infoPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/info/v1"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
)

type InfoPolicy interface {
	GetAll(ctx context.Context) ([]*model.Info, error)
	Create(ctx context.Context, info model.Info) (model.Info, error)
	GetOne(ctx context.Context, albumID string) (model.Info, error)
}

type server struct {
	infoPolicy InfoPolicy
	infoPb.UnimplementedInfoServiceServer
}

func NewServer(policy InfoPolicy, s infoPb.UnimplementedInfoServiceServer) *server {
	return &server{
		infoPolicy:                     policy,
		UnimplementedInfoServiceServer: s,
	}
}
