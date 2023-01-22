package dao

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	dbFIlter "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumDAO struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
}

const (
	table = "album"
)

func NewProductStorage(client PostgreSQLClient) *AlbumDAO {
	return &AlbumDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (a *AlbumDAO) All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]*AlbumDAO, error) {
	dbFIlter.NewFilters(filtering)
	return nil, nil
}
