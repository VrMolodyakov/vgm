package dao

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/jackc/pgx"
)

type professionDAO struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "profession"
)

func NewProfessionStorage(client db.PostgreSQLClient) *professionDAO {
	return &professionDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (p *professionDAO) Create(ctx context.Context, profession string) (ProfessionStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	sql, args, err := p.queryBuilder.
		Insert(table).
		Columns("profession_title").
		Values(profession).
		Suffix(`
				RETURNING profession_id,profession_title
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return ProfessionStorage{}, err
	}

	var pS ProfessionStorage
	if QueryRow := p.client.QueryRow(ctx, sql, args...).
		Scan(
			&pS.ID,
			&pS.Title); QueryRow != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			QueryRow = db.ErrDoQuery(errors.New("profession was not created. 0 rows were affected"))
		} else {
			QueryRow = db.ErrDoQuery(QueryRow)
		}
		logger.Error(QueryRow.Error())
		return ProfessionStorage{}, QueryRow
	}
	return pS, nil
}

func (p *professionDAO) GetOne(ctx context.Context, prof string) (ProfessionStorage, error) {
	logger := logging.LoggerFromContext(ctx)

	query := p.queryBuilder.
		Select(
			"profession_id",
			"profession_title").
		From(table).
		Where(sq.Eq{"profession_title": prof})

	sql, args, err := query.ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return ProfessionStorage{}, err
	}

	var professionStorage ProfessionStorage
	err = p.client.QueryRow(ctx, sql, args...).
		Scan(
			&professionStorage.ID,
			&professionStorage.Title)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err.Error())
		return ProfessionStorage{}, err
	}
	return professionStorage, nil
}
