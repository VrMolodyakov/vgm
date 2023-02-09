package info

import (
	"context"

	infoPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/info/v1"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
)

type InfoPolicy interface {
	Create(ctx context.Context, Info model.Info) (model.Info, error)
	GetOne(ctx context.Context, albumID string) (model.Info, error)
	Update(ctx context.Context, info model.Info) error
	Delete(ctx context.Context, id string) error
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
