package dao

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
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

func (c *creditDAO) Create(ctx context.Context, credits []model.Credit) error {
	logger := logging.LoggerFromContext(ctx)
	insertState := c.queryBuilder.Insert(table).Columns("album_id", "person_id", "credit_role")
	for _, track := range credits {
		insertState = insertState.Values(track.AlbumID, track.Title, track.Duration)
	}
	sql, args, err := insertState.ToSql()
	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return err
	}
	if exec, execErr := t.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = db.ErrDoQuery(execErr)
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Insert() {
		execErr = db.ErrDoQuery(errors.New("tracklist was not created. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil
}

func (c *creditDAO) GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error) {
	logger := logging.LoggerFromContext(ctx)

	query := c.queryBuilder.
		Select(
			"profession_title",
			"first_name",
			"last_name").
		From(table).
		Join("person using (person_id)").
		Join("profession using (profession_id)").
		Join("album using (album_id)").
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

func (c *creditDAO) Delete(ctx context.Context, albumID string) error {
	logger := logging.LoggerFromContext(ctx)
	sql, args, buildErr := c.queryBuilder.
		Delete(table).
		Where(sq.Eq{"album_info_id": albumID}).
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
		execErr = db.ErrDoQuery(errors.New("album was not deleted. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil

}
