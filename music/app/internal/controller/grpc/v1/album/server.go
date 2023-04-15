package album

import (
	"context"

	musicPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumPolicy interface {
	GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumView, error)
	GetOne(ctx context.Context, albumID string) (model.FullAlbum, error)
	Create(ctx context.Context, album model.Album) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.AlbumView) error
}

type server struct {
	albumPolicy AlbumPolicy
	musicPb.UnimplementedMusicServiceServer
}

func NewServer(
	albumPolicy AlbumPolicy,
	s musicPb.UnimplementedMusicServiceServer) *server {

	return &server{
		albumPolicy:                     albumPolicy,
		UnimplementedMusicServiceServer: s,
	}
}
