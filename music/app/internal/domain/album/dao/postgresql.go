package dao

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
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
	client       db.PostgreSQLClient
}

const (
	table = "album"
)

func NewAlbumStorage(client db.PostgreSQLClient) *AlbumDAO {
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

	query = filter.Filter(query, "")
	query = sort.Sort(query, "")

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

func (a *AlbumDAO) Create(ctx context.Context, album model.Album) (AlbumStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	albumStorageMap := ToStorageMap(album)
	sql, args, err := a.queryBuilder.
		Insert(table).
		SetMap(albumStorageMap).
		Suffix("RETURNING album_id ,title, released_at ,created_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return AlbumStorage{}, err
	}

	var albumStorage AlbumStorage

	if QueryRow := a.client.QueryRow(ctx, sql, args...).
		Scan(&albumStorage.ID,
			&albumStorage.Title,
			&albumStorage.ReleasedAt,
			&albumStorage.CreatedAt); QueryRow != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			QueryRow = db.ErrDoQuery(errors.New("album was not created. 0 rows were affected"))
		} else {
			QueryRow = db.ErrDoQuery(QueryRow)
		}
		logger.Error(QueryRow.Error())
		return AlbumStorage{}, QueryRow
	}
	return albumStorage, nil
}

func (a *AlbumDAO) Delete(ctx context.Context, id string) error {
	logger := logging.LoggerFromContext(ctx)
	sql, args, buildErr := a.queryBuilder.
		Delete(table).
		Where(sq.Eq{"album_id": id}).
		ToSql()

	logger.Infow(table, sql, args)

	if buildErr != nil {
		buildErr = db.ErrCreateQuery(buildErr)
		logger.Error(buildErr.Error())
		return buildErr
	}

	if exec, execErr := a.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = db.ErrDoQuery(execErr)
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Delete() {
		execErr = db.ErrDoQuery(errors.New("album was not deleted. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil

}

func (s *AlbumDAO) Update(ctx context.Context, album model.Album) error {
	logger := logging.LoggerFromContext(ctx)
	albumStorageMap := ToUpdateStorageMap(album)
	logger.Info("STORAGE MAP")
	logger.Sugar().Info(albumStorageMap)
	sql, args, buildErr := s.queryBuilder.
		Update(table).
		SetMap(albumStorageMap).
		Where(sq.Eq{"album_id": album.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)

	if buildErr != nil {
		buildErr = db.ErrCreateQuery(buildErr)
		logger.Error(buildErr.Error())
		return buildErr
	}

	if exec, execErr := s.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = db.ErrDoQuery(execErr)
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Update() {
		execErr = db.ErrDoQuery(errors.New("album was not updated. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil
}
