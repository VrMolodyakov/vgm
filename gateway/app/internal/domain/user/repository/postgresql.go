package repository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
)

const (
	table string = "users"
)

type repo struct {
	client       postgresql.PostgreSQLClient
	queryBuilder sq.StatementBuilderType
}

func NewUserRepo(client postgresql.PostgreSQLClient) *repo {
	return &repo{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (r *repo) Create(ctx context.Context, user model.User) error {
	logger := logging.LoggerFromContext(ctx)
	columns := []string{"user_name", "user_email", "user_password", "create_at"}
	nestedSql := r.queryBuilder.
		Select("user_id").
		Prefix("NOT EXISTS(").
		From(table).
		Where(sq.Eq{"user_name": user.Username}).
		Suffix(")")

	notExistSelect := r.queryBuilder.Select(columns...).From(table).Where(nestedSql)
	sql, args, err := r.queryBuilder.Insert(table).Columns(columns...).Select(notExistSelect).ToSql()

	logger.Logger.Sugar().Infow(table, sql, args)

	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return err
	}
	if exec, execErr := r.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = db.ErrDoQuery(execErr)
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Insert() {
		execErr = db.ErrDoQuery(errors.New("user was not created. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}
	return nil
}

func (r *repo) GetOne(ctx context.Context, username string) (model.User, error) {
	logger := logging.LoggerFromContext(ctx)
	query := r.queryBuilder.
		Select(
			"user_id",
			"user_name",
			"user_email",
			"user_password",
			"create_at").
		From(table).
		Where(sq.Eq{"user_name": username})

	sql, args, err := query.ToSql()
	logger.Logger.Sugar().Infow(table, sql, args)

	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return model.User{}, err
	}

	var user model.User
	err = r.client.QueryRow(ctx, sql, args...).
		Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreateAt)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err.Error())
		return model.User{}, err
	}
	return user, nil
}

func (r *repo) Delete(ctx context.Context, username string) error {
	logger := logging.LoggerFromContext(ctx)
	sql, args, buildErr := r.queryBuilder.
		Delete(table).
		Where(sq.Eq{"user_name": username}).
		ToSql()

	logger.Logger.Sugar().Infow(table, sql, args)

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

func (r *repo) Update(ctx context.Context, user model.User) error {
	logger := logging.LoggerFromContext(ctx)
	infoStorageMap := toUpdateStorageMap(user)

	sql, args, buildErr := r.queryBuilder.
		Update(table).
		SetMap(infoStorageMap).
		Where(sq.Eq{"user_name": user.Username}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Logger.Sugar().Infow(table, sql, args)

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
