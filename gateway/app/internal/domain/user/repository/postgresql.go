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
	columns := []string{"user_name", "user_mail", "user_password", "create_at"}
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

func (u *repo) Find(ctx context.Context, username string) (model.User, error) {
	sql := `SELECT u_id,u_name,u_password,create_at FROM users WHERE u_name = $1`
	var user model.User
	err := u.client.QueryRow(ctx, sql, username).Scan(&user.Id, &user.Username, &user.Password, &user.CreateAt)
	if err != nil {

		return model.User{}, err
	}
	return user, nil
}

func (u *repo) FindById(ctx context.Context, id int) (model.User, error) {
	sql := `SELECT u_id,u_name,u_password,create_at FROM users WHERE u_id = $1`
	var user model.User
	err := u.client.QueryRow(ctx, sql, id).Scan(&user.Id, &user.Username, &user.Password, &user.CreateAt)
	if err != nil {

		return model.User{}, err
	}
	return user, nil
}
