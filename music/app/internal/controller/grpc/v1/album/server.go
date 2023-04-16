package album

import (
	"context"

	musicPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"

	albumModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	personModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumPolicy interface {
	GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]albumModel.AlbumView, error)
	GetOne(ctx context.Context, albumID string) (albumModel.FullAlbum, error)
	Create(ctx context.Context, album albumModel.Album) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album albumModel.AlbumView) error
}

type PersonPolicy interface {
	GetAll(ctx context.Context, filter filter.Filterable) ([]personModel.Person, error)
	Create(ctx context.Context, P personModel.Person) (personModel.Person, error)
}

type server struct {
	albumPolicy  AlbumPolicy
	personPolicy PersonPolicy
	musicPb.UnimplementedMusicServiceServer
}

func NewServer(
	albumPolicy AlbumPolicy,
	personPolicy PersonPolicy,
	s musicPb.UnimplementedMusicServiceServer) *server {

	return &server{
		albumPolicy:                     albumPolicy,
		personPolicy:                    personPolicy,
		UnimplementedMusicServiceServer: s,
	}
}
