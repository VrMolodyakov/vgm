package reposotory

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	dbFIlter "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/psqltx"
	dbSort "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/sort"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type Album interface {
	psqltx.Transactor
	Tx(ctx context.Context, action func(txRepo Album) error) error
	GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumView, error)
	GetInfo(ctx context.Context, albumID string) (model.AlbumInfo, error)
	Create(ctx context.Context, album model.Album) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.AlbumView) error
}

type repo struct {
	queryBuilder sq.StatementBuilderType
	psqltx.Transactor
}

const (
	table          = "album"
	infoTabe       = "album_info"
	creditTabe     = "credit"
	tracklistTable = "track"
)

func NewAlbumRepository(client db.PostgreSQLClient) Album {
	return &repo{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		Transactor:   psqltx.NewTx(client),
	}
}

func (r *repo) GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumView, error) {
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
		err = errors.NewInternal(db.ErrCreateQuery(err), "create query")
		logger.Error(err.Error())
		return nil, err
	}
	rows, queryErr := r.Conn().Query(ctx, sql, args...)
	if queryErr != nil {
		err = errors.NewInternal(db.ErrCreateQuery(err), "do query")
		logger.Error(err.Error())
		return nil, err
	}

	albums := make([]model.AlbumView, 0)
	for rows.Next() {
		as := AlbumViewStorage{}
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

func (r *repo) Delete(ctx context.Context, id string) error {
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

func (r *repo) DeleteInfo(ctx context.Context, id string) error {
	logger := logging.LoggerFromContext(ctx)
	sql, args, buildErr := r.queryBuilder.
		Delete(table).
		Where(sq.Eq{"album_info_id": id}).
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

func (r *repo) Update(ctx context.Context, album model.AlbumView) error {
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

func (r *repo) UpdateInfo(ctx context.Context, info model.Info) error {
	logger := logging.LoggerFromContext(ctx)
	infoStorageMap := toUpdateStorageMap(&info)

	sql, args, buildErr := r.queryBuilder.
		Update(table).
		SetMap(infoStorageMap).
		Where(sq.Eq{"album_info_id": info.ID}).
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

func (r *repo) GetInfo(ctx context.Context, albumID string) (model.AlbumInfo, error) {
	logger := logging.LoggerFromContext(ctx)
	query := r.queryBuilder.
		Select(
			"album_id",
			"title",
			"released_at",
			"created_at",
			"album_info_id",
			"album_id",
			"catalog_number",
			"full_image_src",
			"small_image_src",
			"barcode",
			"price",
			"currency_code",
			"media_format",
			"classification",
			"publisher").
		From(table).
		Join("album_info USING (album_id)").
		Where(sq.Eq{"album_id": albumID})

	sql, args, err := query.ToSql()
	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return model.AlbumInfo{}, err
	}

	var storage AlbumInfoStorage
	err = r.Conn().QueryRow(ctx, sql, args...).
		Scan(
			&storage.Album.ID,
			&storage.Album.Title,
			&storage.Album.ReleasedAt,
			&storage.Album.CreatedAt,
			&storage.Info.ID,
			&storage.Info.AlbumID,
			&storage.Info.CatalogNumber,
			&storage.Info.FullImageSrc,
			&storage.Info.SmallImageSrc,
			&storage.Info.Barcode,
			&storage.Info.Price,
			&storage.Info.CurrencyCode,
			&storage.Info.MediaFormat,
			&storage.Info.Classification,
			&storage.Info.Publisher)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err.Error())
		return model.AlbumInfo{}, err
	}
	return storage.toModel(), nil
}

func (r *repo) Tx(ctx context.Context, action func(txRepo Album) error) error {
	return r.WithinTransaction(
		ctx,
		func(client db.PostgreSQLClient) psqltx.Transactor { return NewAlbumRepository(client) },
		func(txRepo psqltx.Transactor) error { return action(txRepo.(Album)) },
	)
}

func (r *repo) Create(ctx context.Context, album model.Album) error {
	logger := logging.LoggerFromContext(ctx)
	albumStorageMap := ToStorageMap(album.Album)
	tx, err := r.Conn().Begin(ctx)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	if err != nil {
		err = errors.NewInternal(err, "start tx")
		logger.Error(err.Error())
		return err
	}
	sql, args, err := r.insertMap(table, albumStorageMap)
	logger.Infow(table, sql, args)
	if err != nil {
		err = errors.NewInternal(db.ErrCreateQuery(err), "create query")
		logger.Error(err.Error())
		return err

	}
	err = exec(tx, ctx, sql, args)
	if err != nil {
		return err
	}
	infoStorageMap := ToInfoStorageMap(&album.Info)
	sql, args, err = r.insertMap(infoTabe, infoStorageMap)
	if err != nil {
		err = errors.NewInternal(db.ErrCreateQuery(err), "create query")
		logger.Error(err.Error())
		return err
	}
	err = exec(tx, ctx, sql, args)
	if err != nil {
		return err
	}

	insertState := r.queryBuilder.Insert(creditTabe).Columns("album_id", "person_id", "credit_role")
	for _, credit := range album.Credits {
		insertState = insertState.Values(credit.AlbumID, credit.PersonID, credit.Profession)
	}
	sql, args, err = insertState.ToSql()
	logger.Infow(table, sql, args)
	if err != nil {
		err = errors.NewInternal(db.ErrCreateQuery(err), "create query")
		logger.Error(err.Error())
		return err
	}
	err = exec(tx, ctx, sql, args)
	if err != nil {
		return err
	}
	insertState = r.queryBuilder.Insert(tracklistTable).Columns("album_id", "title", "duration")
	for _, track := range album.Tracklist {
		insertState = insertState.Values(track.AlbumID, track.Title, track.Duration)
	}
	sql, args, err = insertState.ToSql()
	logger.Infow(table, sql, args)
	if err != nil {
		err = errors.NewInternal(db.ErrCreateQuery(err), "create query")
		logger.Error(err.Error())
		return err
	}
	err = exec(tx, ctx, sql, args)
	if err != nil {
		return err
	}

	return nil

}

func (r *repo) insertMap(table string, storageMap map[string]interface{}) (string, []interface{}, error) {
	return r.queryBuilder.
		Insert(table).
		SetMap(storageMap).
		PlaceholderFormat(sq.Dollar).
		ToSql()
}

func exec(client db.PostgreSQLClient, ctx context.Context, sql string, args []interface{}) error {
	logger := logging.LoggerFromContext(ctx)
	if exec, execErr := client.Exec(ctx, sql, args...); execErr != nil {
		execErr = errors.NewInternal(db.ErrDoQuery(execErr), "exec query")
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Insert() {
		execErr = db.ErrDoQuery(errors.New("album was not created. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}
	return nil
}
