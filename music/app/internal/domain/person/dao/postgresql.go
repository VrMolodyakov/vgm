package dao

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	dbFIlter "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/jackc/pgx"
)

type personDAO struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "person"
)

func NewPersonStorage(client db.PostgreSQLClient) *personDAO {
	return &personDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (p *personDAO) Create(ctx context.Context, person model.Person) (PersonStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	personStorageMap := toStorageMap(person)
	sql, args, err := p.queryBuilder.
		Insert(table).
		SetMap(personStorageMap).
		Suffix(`
				RETURNING person_id,first_name,last_name,birth_date
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return PersonStorage{}, err
	}

	var personStorage PersonStorage
	if QueryRow := p.client.QueryRow(ctx, sql, args...).
		Scan(
			&personStorage.ID,
			&personStorage.FirstName,
			&personStorage.LastName,
			&personStorage.BirthDate); QueryRow != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			QueryRow = db.ErrDoQuery(errors.New("person was not created. 0 rows were affected"))
		} else {
			QueryRow = db.ErrDoQuery(QueryRow)
		}
		logger.Error(QueryRow.Error())
		return PersonStorage{}, QueryRow
	}
	return personStorage, nil
}

func (p *personDAO) GetAll(ctx context.Context, filtering filter.Filterable) ([]PersonStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	filter := dbFIlter.NewFilters(filtering)
	query := p.queryBuilder.
		Select("person_id", "first_name", "last_name", "birth_date").
		From(table)

	query = filter.Filter(query, "")

	sql, args, err := query.ToSql()
	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return nil, err
	}
	rows, queryErr := p.client.Query(ctx, sql, args...)
	if queryErr != nil {
		err := db.ErrDoQuery(queryErr)
		logger.Error(err.Error())
		return nil, err
	}

	persons := make([]PersonStorage, 0)
	for rows.Next() {
		p := PersonStorage{}
		if queryErr = rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.BirthDate,
		); queryErr != nil {
			queryErr = db.ErrScan(queryErr)
			logger.Error(queryErr.Error())
			return nil, queryErr
		}

		persons = append(persons, p)
	}

	return persons, nil
}

func (p *personDAO) GetOne(ctx context.Context, personID string) (PersonStorage, error) {
	logger := logging.LoggerFromContext(ctx)

	query := p.queryBuilder.
		Select(
			"person_id",
			"first_name",
			"last_name",
			"birth_date").
		From(table).
		Where(sq.Eq{"person_id": personID})

	sql, args, err := query.ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return PersonStorage{}, err
	}

	var pS PersonStorage
	err = p.client.QueryRow(ctx, sql, args...).
		Scan(
			&pS.ID,
			&pS.FirstName,
			&pS.LastName,
			&pS.BirthDate)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err.Error())
		return PersonStorage{}, err
	}
	return pS, nil
}
