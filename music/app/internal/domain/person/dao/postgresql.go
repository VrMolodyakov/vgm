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
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
	"github.com/jackc/pgx"
)

type PersonDAO struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "person"
)

func NewPersonStorage(client db.PostgreSQLClient) *PersonDAO {
	return &PersonDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (a *PersonDAO) Create(ctx context.Context, person model.Person) (PersonStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	personStorageMap := toStorageMap(person)
	sql, args, err := a.queryBuilder.
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
	if QueryRow := a.client.QueryRow(ctx, sql, args...).
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

func (a *PersonDAO) GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]PersonStorage, error) {
	logger := logging.LoggerFromContext(ctx)
	filter := dbFIlter.NewFilters(filtering)
	query := a.queryBuilder.
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
	rows, queryErr := a.client.Query(ctx, sql, args...)
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
