package reposotory

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	dbFIlter "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/jackc/pgx"
)

type repo struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "person"
)

func NewPersonStorage(client db.PostgreSQLClient) *repo {
	return &repo{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (r *repo) Create(ctx context.Context, person model.Person) (model.Person, error) {
	logger := logging.LoggerFromContext(ctx)
	personStorageMap := toStorageMap(person)
	sql, args, err := r.queryBuilder.
		Insert(table).
		SetMap(personStorageMap).
		Suffix(`
				RETURNING person_id,first_name,last_name,birth_date
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = errors.NewInternal(err, "create sql")
		logger.Error(err.Error())
		return model.Person{}, err
	}

	var storage personStorage
	if QueryRow := r.client.QueryRow(ctx, sql, args...).
		Scan(
			&storage.ID,
			&storage.FirstName,
			&storage.LastName,
			&storage.BirthDate); QueryRow != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			QueryRow = errors.NewInternal(err, "person was not created. 0 rows were affected")
		} else {
			QueryRow = errors.NewInternal(err, "do query")
		}
		logger.Error(QueryRow.Error())
		return model.Person{}, QueryRow
	}
	return storage.toModel(), nil
}

func (r *repo) GetAll(ctx context.Context, filtering filter.Filterable) ([]model.Person, error) {
	logger := logging.LoggerFromContext(ctx)
	filter := dbFIlter.NewFilters(filtering)
	query := r.queryBuilder.
		Select("person_id", "first_name", "last_name", "birth_date").
		From(table)

	query = filter.Filter(query, "")

	sql, args, err := query.ToSql()
	logger.Infow(table, sql, args)
	if err != nil {
		err = errors.NewInternal(err, "create query")
		logger.Error(err.Error())
		return nil, err
	}
	rows, queryErr := r.client.Query(ctx, sql, args...)
	if queryErr != nil {
		err := errors.NewInternal(queryErr, "do query")
		logger.Error(err.Error())
		return nil, err
	}

	persons := make([]model.Person, 0)
	for rows.Next() {
		p := personStorage{}
		if queryErr = rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.BirthDate,
		); queryErr != nil {
			queryErr = errors.NewInternal(queryErr, "scan query")
			logger.Error(queryErr.Error())
			return nil, queryErr
		}

		persons = append(persons, p.toModel())
	}

	return persons, nil
}

func (r *repo) GetOne(ctx context.Context, personID string) (model.Person, error) {
	logger := logging.LoggerFromContext(ctx)

	query := r.queryBuilder.
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
		return model.Person{}, err
	}

	var storage personStorage
	err = r.client.QueryRow(ctx, sql, args...).
		Scan(
			&storage.ID,
			&storage.FirstName,
			&storage.LastName,
			&storage.BirthDate)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err.Error())
		return model.Person{}, err
	}
	return storage.toModel(), nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	logger := logging.LoggerFromContext(ctx)
	sql, args, buildErr := r.queryBuilder.
		Delete(table).
		Where(sq.Eq{"person_id": id}).
		ToSql()

	logger.Infow(table, sql, args)

	if buildErr != nil {
		buildErr = db.ErrCreateQuery(buildErr)
		logger.Error(buildErr.Error())
		return buildErr
	}

	if exec, execErr := r.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = db.ErrDoQuery(execErr)
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Delete() {
		execErr = db.ErrDoQuery(errors.New("person was not deleted. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil

}

func (r *repo) Update(ctx context.Context, person model.Person) error {
	logger := logging.LoggerFromContext(ctx)
	infoStorageMap := toUpdateStorageMap(person)

	sql, args, buildErr := r.queryBuilder.
		Update(table).
		SetMap(infoStorageMap).
		Where(sq.Eq{"person_id ": person.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Infow(table, sql, args)

	if buildErr != nil {
		buildErr = db.ErrCreateQuery(buildErr)
		logger.Error(buildErr.Error())
		return buildErr
	}

	if exec, execErr := r.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = db.ErrDoQuery(execErr)
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Update() {
		execErr = db.ErrDoQuery(errors.New("person was not updated. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil
}
