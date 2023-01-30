package dao

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/jackc/pgx"
)

type InfoDAO struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "album_info"
)

func NewInfoStorage(client db.PostgreSQLClient) *InfoDAO {
	return &InfoDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (a *InfoDAO) All(ctx context.Context) ([]InfoStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	query := a.queryBuilder.
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
		From(table)

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

	albums := make([]InfoStorage, 0)
	for rows.Next() {
		as := InfoStorage{}
		if queryErr = rows.Scan(
			&as.ID,
			&as.AlbumID,
			&as.CatalogNumber,
			&as.ImageSrc,
			&as.Barcode,
			&as.Price,
			&as.CurrencyCode,
			&as.MediaFormat,
			&as.Classification,
			&as.Publisher,
		); queryErr != nil {
			queryErr = db.ErrScan(queryErr)
			logger.Error(queryErr.Error())
			return nil, queryErr
		}

		albums = append(albums, as)
	}

	return albums, nil
}

func (a *InfoDAO) Create(ctx context.Context, info model.Info) (InfoStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	infoStorageMap := ToStorageMap(info)
	sql, args, err := a.queryBuilder.
		Insert(table).
		SetMap(infoStorageMap).
		Suffix(`
				RETURNING album_info_id,album_id,catalog_number,image_srs,
				barcode,
				price,
				currency_code,
				media_format,
				classification,
				publisher`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return InfoStorage{}, err
	}

	var infoStorage InfoStorage
	if QueryRow := a.client.QueryRow(ctx, sql, args...).
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
			&infoStorage.Publisher); QueryRow != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			QueryRow = db.ErrDoQuery(errors.New("album info was not created. 0 rows were affected"))
		} else {
			QueryRow = db.ErrDoQuery(QueryRow)
		}
		logger.Error(QueryRow.Error())
		return InfoStorage{}, QueryRow
	}
	return infoStorage, nil
}
