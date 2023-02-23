package reposotory

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type repo struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "credit"
)

func NewCreditRepo(client db.PostgreSQLClient) *repo {
	return &repo{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (c *repo) GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error) {
	logger := logging.LoggerFromContext(ctx)

	query := c.queryBuilder.
		Select(
			"credit_role",
			"first_name",
			"last_name").
		From(table).
		Join("person using (person_id)").
		Where(sq.Eq{"album_id": albumID})

	sql, args, err := query.ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return nil, err
	}

	rows, queryErr := c.client.Query(ctx, sql, args...)
	if queryErr != nil {
		queryErr = db.ErrDoQuery(queryErr)
		logger.Error(queryErr.Error())
		return nil, queryErr
	}
	creditInfo := make([]model.CreditInfo, 0)
	for rows.Next() {
		storage := CreditInfoStorage{}
		if queryErr = rows.Scan(
			&storage.Profession,
			&storage.FirstName,
			&storage.LastName,
		); queryErr != nil {
			queryErr = db.ErrScan(queryErr)
			logger.Error(queryErr.Error())
			return nil, queryErr
		}
		creditInfo = append(creditInfo, storage.toModel())

	}
	return creditInfo, nil
}

func (c *repo) Delete(ctx context.Context, albumID string) error {
	logger := logging.LoggerFromContext(ctx)
	sql, args, buildErr := c.queryBuilder.
		Delete(table).
		Where(sq.Eq{"album_id": albumID}).
		ToSql()

	logger.Infow(table, sql, args)

	if buildErr != nil {
		buildErr = db.ErrCreateQuery(buildErr)
		logger.Error(buildErr.Error())
		return buildErr
	}

	if exec, execErr := c.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = db.ErrDoQuery(execErr)
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Delete() {
		execErr = db.ErrDoQuery(errors.New("credit was not deleted. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil

}

func (r *repo) Update(ctx context.Context, albumId string, role string) error {
	logger := logging.LoggerFromContext(ctx)

	sql, args, buildErr := r.queryBuilder.
		Update(table).
		Set("credit_role", role).
		Where(sq.Eq{"album_id": albumId}).
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
		execErr = db.ErrDoQuery(errors.New("credit was not updated. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil
}
