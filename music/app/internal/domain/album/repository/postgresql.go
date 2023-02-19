package reposotory

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	dbFIlter "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/psqltx"
	dbSort "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/sort"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type Album interface {
	psqltx.Transactor
	Tx(ctx context.Context, action func(txRepo Album) error) error
	GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumView, error)
	GetOne(ctx context.Context, albumID string) (model.AlbumView, error)
	Create(ctx context.Context, album model.AlbumView) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.AlbumView) error
}

type repository struct {
	queryBuilder sq.StatementBuilderType
	psqltx.Transactor
}

const (
	table = "album"
)

func NewAlbumRepository(client db.PostgreSQLClient) Album {
	return &repository{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		Transactor:   psqltx.NewTx(client),
	}
}

func (r *repository) GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumView, error) {
	logger := logging.LoggerFromContext(ctx)
	filter := dbFIlter.NewFilters(filtering)
	sort := dbSort.NewSortOptions(sorting)
	query := r.queryBuilder.
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
	rows, queryErr := r.Conn().Query(ctx, sql, args...)
	if queryErr != nil {
		err := db.ErrDoQuery(queryErr)
		logger.Error(err.Error())
		return nil, err
	}

	albums := make([]model.AlbumView, 0)
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
		albums = append(albums, as.toModel())
	}

	return albums, nil
}

func (r *repository) Create(ctx context.Context, album model.AlbumView) error {
	logger := logging.LoggerFromContext(ctx)
	albumStorageMap := toStorageMap(album)
	sql, args, err := r.queryBuilder.
		Insert(table).
		SetMap(albumStorageMap).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return err
	}
	if exec, execErr := r.Conn().Exec(ctx, sql, args...); execErr != nil {
		execErr = db.ErrDoQuery(execErr)
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Insert() {
		execErr = db.ErrDoQuery(errors.New("album was not created. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	logger := logging.LoggerFromContext(ctx)
	sql, args, buildErr := r.queryBuilder.
		Delete(table).
		Where(sq.Eq{"album_id": id}).
		ToSql()

	logger.Infow(table, sql, args)

	if buildErr != nil {
		buildErr = db.ErrCreateQuery(buildErr)
		logger.Error(buildErr.Error())
		return buildErr
	}

	if exec, execErr := r.Conn().Exec(ctx, sql, args...); execErr != nil {
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

func (r *repository) Update(ctx context.Context, album model.AlbumView) error {
	logger := logging.LoggerFromContext(ctx)
	albumStorageMap := ToUpdateStorageMap(album)
	sql, args, buildErr := r.queryBuilder.
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

	if exec, execErr := r.Conn().Exec(ctx, sql, args...); execErr != nil {
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

func (r *repository) GetOne(ctx context.Context, albumID string) (model.AlbumView, error) {
	logger := logging.LoggerFromContext(ctx)
	query := r.queryBuilder.
		Select("album_id", "title", "released_at", "created_at").
		From(table).
		Where(sq.Eq{"album_id": albumID})

	sql, args, err := query.ToSql()
	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return model.AlbumView{}, err
	}

	var storage AlbumStorage
	err = r.Conn().QueryRow(ctx, sql, args...).
		Scan(
			&storage.ID,
			&storage.Title,
			&storage.ReleasedAt,
			&storage.CreatedAt)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err.Error())
		return model.AlbumView{}, err
	}
	return storage.toModel(), nil
}

func (r *repository) Tx(ctx context.Context, action func(txRepo Album) error) error {
	return r.WithinTransaction(
		ctx,
		func(client db.PostgreSQLClient) psqltx.Transactor { return NewAlbumRepository(client) },
		func(txRepo psqltx.Transactor) error { return action(txRepo.(Album)) },
	)
}
