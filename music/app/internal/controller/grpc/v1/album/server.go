package album

import (
	"context"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumPolicy interface {
	All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.Album, error)
	CreateProduct(ctx context.Context, product model.Album) (model.Album, error)
}

type server struct {
	albumPolicy AlbumPolicy
	albumPb.UnimplementedAlbumServiceServer
}

func NewServer(policy AlbumPolicy, s albumPb.UnimplementedAlbumServiceServer) *server {
	return &server{
		albumPolicy:                     policy,
		UnimplementedAlbumServiceServer: s,
	}
}
