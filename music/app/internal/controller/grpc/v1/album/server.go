package album

import (
	"context"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"

	albumModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	infoModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	trackModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumPolicy interface {
	GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]albumModel.Album, error)
	Create(ctx context.Context, album albumModel.Album) (albumModel.Album, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album albumModel.Album) error
}

type InfoPolicy interface {
	Create(ctx context.Context, info infoModel.Info) (infoModel.Info, error)
}

type TracklistPolicy interface {
	Create(ctx context.Context, tracklist []trackModel.Track) error
}

type server struct {
	albumPolicy AlbumPolicy
	infoPolicy  InfoPolicy
	trackPolicy TracklistPolicy
	albumPb.UnimplementedAlbumServiceServer
}

func NewServer(
	albumPolicy AlbumPolicy,
	infoPolicy InfoPolicy,
	tracklistPolicy TracklistPolicy,
	s albumPb.UnimplementedAlbumServiceServer) *server {

	return &server{
		albumPolicy:                     albumPolicy,
		infoPolicy:                      infoPolicy,
		trackPolicy:                     tracklistPolicy,
		UnimplementedAlbumServiceServer: s,
	}
}
