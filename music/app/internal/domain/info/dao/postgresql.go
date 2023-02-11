package dao

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type infoDAO struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "album_info"
)

func NewInfoStorage(client db.PostgreSQLClient) *infoDAO {
	return &infoDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (i *infoDAO) Create(ctx context.Context, info model.Info) error {
	logger := logging.LoggerFromContext(ctx)
	infoStorageMap := toStorageMap(&info)
	sql, args, err := i.queryBuilder.
		Insert(table).
		SetMap(infoStorageMap).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return err
	}
	if exec, execErr := i.client.Exec(ctx, sql, args...); execErr != nil {
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

func (i *infoDAO) GetOne(ctx context.Context, albumID string) (model.Info, error) {
	logger := logging.LoggerFromContext(ctx)

	query := i.queryBuilder.
		Select(
			"album_info_id",
			"album_id",
			"catalog_number",
			"image_srs",
			"barcode",
			"price",
			"currency_code",
			"media_format",
			"classification",
			"publisher").
		From(table).
		Where(sq.Eq{"album_id": albumID})

	sql, args, err := query.ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return model.Info{}, err
	}

	var infoStorage InfoStorage
	err = i.client.QueryRow(ctx, sql, args...).
		Scan(
			&infoStorage.ID,
			&infoStorage.AlbumID,
			&infoStorage.CatalogNumber,
			&infoStorage.ImageSrc,
			&infoStorage.Barcode,
			&infoStorage.Price,
			&infoStorage.CurrencyCode,
			&infoStorage.MediaFormat,
			&infoStorage.Classification,
			&infoStorage.Publisher)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err.Error())
		return model.Info{}, err
	}
	return infoStorage.toModel(), nil
}

func (i *infoDAO) Delete(ctx context.Context, id string) error {
	logger := logging.LoggerFromContext(ctx)
	sql, args, buildErr := i.queryBuilder.
		Delete(table).
		Where(sq.Eq{"album_info_id": id}).
		ToSql()

	logger.Infow(table, sql, args)

	if buildErr != nil {
		buildErr = db.ErrCreateQuery(buildErr)
		logger.Error(buildErr.Error())
		return buildErr
	}

	if exec, execErr := i.client.Exec(ctx, sql, args...); execErr != nil {
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

func (i *infoDAO) Update(ctx context.Context, info model.Info) error {
	logger := logging.LoggerFromContext(ctx)
	infoStorageMap := toUpdateStorageMap(&info)

	sql, args, buildErr := i.queryBuilder.
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

	if exec, execErr := i.client.Exec(ctx, sql, args...); execErr != nil {
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
