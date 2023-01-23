package dao

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
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

func (a *AlbumDAO) All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]AlbumStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	filter := dbFIlter.NewFilters(filtering)
	sort := dbSort.NewSortOptions(sorting)

	query := a.queryBuilder.
		Select("album_id", "title", "create_at").
		From(table)

	filter.Filter(query, "")
	sort.Sort(query, "")

	sql, args, err := query.ToSql()
	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return nil, err
	}
	rows, err := a.client.Query(ctx, sql, args...)
	if err != nil {
		err := db.ErrDoQuery(err)
		logger.Error(err.Error())
		return nil, err
	}

	albums := make([]AlbumStorage, 0)
	for rows.Next() {
		as := AlbumStorage{}
		if err = rows.Scan(
			&as.ID,
			&as.Title,
			&as.CreateAt,
		); err != nil {
			err = db.ErrScan(err)
			logger.Error(err.Error())
			return nil, err
		}

		albums = append(albums, as)
	}

	return albums, nil
}

func (a *AlbumDAO) Create(ctx context.Context, m map[string]interface{}) error {
	logger := logging.LoggerFromContext(ctx)
	sql, args, err := a.queryBuilder.
		Insert(table).
		SetMap(m).
		PlaceholderFormat(sq.Dollar).ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return err
	}

	if exec, execErr := a.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = db.ErrDoQuery(execErr)
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Insert() {
		execErr = db.ErrDoQuery(errors.New("product was not created. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil
}
