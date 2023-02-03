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
	table = "Credit"
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
				RETURNING Credit_id,Credit_title
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
			QueryRow = db.ErrDoQuery(errors.New("Credit was not created. 0 rows were affected"))
		} else {
			QueryRow = db.ErrDoQuery(QueryRow)
		}
		logger.Error(QueryRow.Error())
		return CreditStorage{}, QueryRow
	}
	return cS, nil
}

func (p *creditDAO) GetOne(ctx context.Context, profID string) (CreditStorage, error) {
	logger := logging.LoggerFromContext(ctx)

	query := p.queryBuilder.
		Select(
			"Credit_id",
			"Credit_title").
		From(table).
		Where(sq.Eq{"Credit_id": profID})

	sql, args, err := query.ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return CreditStorage{}, err
	}

	var cS CreditStorage
	err = p.client.QueryRow(ctx, sql, args...).
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
