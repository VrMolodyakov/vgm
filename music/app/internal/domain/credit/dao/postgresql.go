package dao

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/jackc/pgx"
)

type creditDAO struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "credit"
)

func NewCreditStorage(client db.PostgreSQLClient) *creditDAO {
	return &creditDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (c *creditDAO) Create(ctx context.Context, Credit model.Credit) (CreditStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	CreditStorageMap := toStorageMap(Credit)
	sql, args, err := c.queryBuilder.
		Insert(table).
		SetMap(CreditStorageMap).
		Suffix(`
				RETURNING album_id,profession_id,person_id
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return CreditStorage{}, err
	}

	var cS CreditStorage
	if QueryRow := c.client.QueryRow(ctx, sql, args...).
		Scan(
			&cS.AlbumID,
			&cS.ProfessionID,
			&cS.PersonID); QueryRow != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			QueryRow = db.ErrDoQuery(errors.New("credit was not created. 0 rows were affected"))
		} else {
			QueryRow = db.ErrDoQuery(QueryRow)
		}
		logger.Error(QueryRow.Error())
		return CreditStorage{}, QueryRow
	}
	return cS, nil
}

func (c *creditDAO) GetOne(ctx context.Context, albumID string) (CreditStorage, error) {
	logger := logging.LoggerFromContext(ctx)

	query := c.queryBuilder.
		Select(
			"album_id",
			"profession_id",
			"person_id").
		From(table).
		Where(sq.Eq{"album_id": albumID})

	sql, args, err := query.ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return CreditStorage{}, err
	}

	var cS CreditStorage
	err = c.client.QueryRow(ctx, sql, args...).
		Scan(
			&cS.AlbumID,
			&cS.ProfessionID,
			&cS.PersonID)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err.Error())
		return CreditStorage{}, err
	}
	return cS, nil
}
