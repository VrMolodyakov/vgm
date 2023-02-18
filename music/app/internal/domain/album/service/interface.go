package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	creditModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	infoModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	trackModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
	"github.com/jackc/pgx/v4"
)

type AlbumRepo interface {
	GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumView, error)
	GetOne(ctx context.Context, albumID string) (model.AlbumView, error)
	Create(ctx context.Context, album model.AlbumView) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.AlbumView) error
	TX(action func(txRepo AlbumRepo) error) error
}

type CreditRepo interface {
	Create(ctx context.Context, credits []creditModel.Credit) error
}

type InfoRepo interface {
	Create(ctx context.Context, info infoModel.Info) error
}

type TrackRepo interface {
	Create(ctx context.Context, tracklist []trackModel.Track) error
}

type Transactor interface {
	WithinTransaction(ctx context.Context, isoLevel pgx.TxOptions, tFunc func(ctx context.Context) error) error
}
