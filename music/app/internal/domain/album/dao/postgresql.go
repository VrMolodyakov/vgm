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
	"github.com/jackc/pgx"
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
		Select("album_id", "title", "released_at", "created_at").
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
	rows, queryErr := a.client.Query(ctx, sql, args...)
	if queryErr != nil {
		err := db.ErrDoQuery(queryErr)
		logger.Error(err.Error())
		return nil, err
	}

	albums := make([]AlbumStorage, 0)
	for rows.Next() {
		as := AlbumStorage{}
		if queryErr = rows.Scan(
			&as.ID,
			&as.Title,
			&as.ReleasedAt,
			&as.CreatedAt,
		); queryErr != nil {
			queryErr = db.ErrScan(queryErr)
			logger.Error(queryErr.Error())
			return nil, queryErr
		}

		albums = append(albums, as)
	}

	return albums, nil
}

func (a *AlbumDAO) Create(ctx context.Context, m map[string]interface{}) (AlbumStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	sql, args, err := a.queryBuilder.
		Insert(table).
		SetMap(m).
		Suffix("RETURNING album_id ,title, released_at ,created_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return AlbumStorage{}, err
	}

	var album AlbumStorage
	if QueryRow := a.client.QueryRow(ctx, sql, args...).
		Scan(&album.ID,
			&album.Title,
			&album.ReleasedAt,
			&album.CreatedAt); QueryRow != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			QueryRow = db.ErrDoQuery(errors.New("album was not created. 0 rows were affected"))
		} else {
			QueryRow = db.ErrDoQuery(QueryRow)
		}
		logger.Error(QueryRow.Error())
		return AlbumStorage{}, QueryRow
	}

	return album, nil
}
