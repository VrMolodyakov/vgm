package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	creditModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	trackModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumService interface {
	GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumPreview, error)
	GetOne(ctx context.Context, albumID string) (model.AlbumInfo, error)
	GetRandom(ctx context.Context, count uint64) ([]string, error)
	GetDays(ctx context.Context, count uint64) ([]int64, error)
	Create(ctx context.Context, album model.Album) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.AlbumView) error
}

type TrackService interface {
	GetAll(ctx context.Context, albumID string) ([]trackModel.Track, error)
}

type CreditService interface {
	GetAll(ctx context.Context, albumID string) ([]creditModel.CreditInfo, error)
}
