package dao

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	dbFIlter "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/filter"
	dbSort "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/sort"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
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
	logger := logging.LoggerFromContext(ctx)
	filter := dbFIlter.NewFilters(filtering)
	sort := dbSort.NewSortOptions(sorting)
	query := a.queryBuilder.
		Select("id").
		From(table)

	query = filter.Filter(query, "")
	query = sort.Sort(query, "")
	sql, args, err := query.ToSql()
	if err != nil {
		logger.Sugar().With(
			sql,
			args,
		).Error(err.Error())
		return nil, err
	}

	return nil, nil
}
