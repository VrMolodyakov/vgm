package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	creditModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	infoModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	trackModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumService interface {
	GetAll(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.AlbumView, error)
	GetOne(ctx context.Context, albumID string) (model.AlbumView, error)
	Create(ctx context.Context, album model.Album) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.AlbumView) error
}

type InfoService interface {
	GetOne(ctx context.Context, albumID string) (infoModel.Info, error)
}

type TrackService interface {
	GetAll(ctx context.Context, albumID string) ([]trackModel.Track, error)
}

type CreditService interface {
	GetAll(ctx context.Context, albumID string) ([]creditModel.CreditInfo, error)
}
