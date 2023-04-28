package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	repository "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/repository"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
	"github.com/jackc/pgx/v5"
)

type AlbumRepo interface {
	Tx(ctx context.Context, action func(txRepo repository.Album) error) error
	GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumView, error)
	GetInfo(ctx context.Context, albumID string) (model.AlbumInfo, error)
	Create(ctx context.Context, album model.Album) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.AlbumView) error
}

type Transactor interface {
	WithinTransaction(ctx context.Context, isoLevel pgx.TxOptions, tFunc func(ctx context.Context) error) error
}
