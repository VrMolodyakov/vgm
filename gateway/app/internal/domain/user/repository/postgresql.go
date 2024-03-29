package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/user/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel"
)

const (
	userTable      string = "users"
	userRolesTable string = "user_roles"
	rolesTable     string = "roles"
)

var (
	tracer = otel.Tracer("user-repo")
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

func (r *repo) Create(ctx context.Context, user model.User) (int, error) {
	ctx, span := tracer.Start(ctx, "repo.Create")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	tx, err := r.client.Begin(ctx)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	if err != nil {
		err = errors.NewInternal(err, "start tx")
		logger.Error(err.Error())
		return -1, err
	}
	sql, args, err := r.queryBuilder.
		Select("user_id").
		From(userTable).
		Where(sq.Eq{"user_name": user.Username}).ToSql()

	logger.Logger.Sugar().Infow(userTable, sql, args)
	if err != nil {
		err = errors.NewInternal(err, "create query")
		logger.Error(err.Error())
		return -1, err
	}
	var userID int
	err = tx.QueryRow(ctx, sql, args...).Scan(&userID)
	if err != nil {
		logger.Error(err.Error())
		if !errors.Is(err, pgx.ErrNoRows) {
			err = errors.NewInternal(err, "executy query")
			return -1, err
		}
		sql, args, err = r.queryBuilder.
			Insert(userTable).
			Columns(
				"user_name",
				"user_email",
				"user_password",
				"create_at").
			Values(user.Username, user.Email, user.Password, time.Now()).
			Suffix("RETURNING user_id").
			ToSql()

		err = tx.QueryRow(ctx, sql, args...).Scan(&userID)
		if err != nil {
			logger.Error(err.Error())
			if errors.Is(err, pgx.ErrNoRows) {
				return -1, err
			}
			err = errors.NewInternal(err, "executy query")
			return -1, err
		}

		sql, args, err = r.queryBuilder.Select("role_id").From(rolesTable).Where(sq.Eq{"role_name": user.Role}).ToSql()
		if err != nil {
			err = errors.NewInternal(err, "create query")
			logger.Error(err.Error())
			return -1, err
		}
		var roleID int
		err = tx.QueryRow(ctx, sql, args...).Scan(&roleID)
		if err != nil {
			err = errors.NewInternal(err, "executy query")
			logger.Error(err.Error())
			return -1, err
		}
		sql, args, err = r.queryBuilder.Insert(userRolesTable).Columns("user_id", "role_id").Values(userID, roleID).ToSql()
		if err != nil {
			err = errors.NewInternal(err, "create query")
			logger.Error(err.Error())
			return -1, err
		}
		if exec, execErr := tx.Exec(ctx, sql, args...); execErr != nil {
			execErr = errors.NewInternal(execErr, "executy query")
			logger.Error(execErr.Error())
			return -1, execErr
		} else if exec.RowsAffected() == 0 || !exec.Insert() {
			execErr = errors.NewInternal(execErr, "user was not created. 0 rows were affected")
			logger.Error(execErr.Error())
			return -1, execErr
		}
		return userID, nil
	}
	return -1, errors.New("user already exists")

}

func (r *repo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	ctx, span := tracer.Start(ctx, "repo.GetByUsername")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	query := r.queryBuilder.
		Select(
			"user_id",
			"user_name",
			"user_email",
			"user_password",
			"create_at",
			"role_name").
		From(userTable).
		Join("user_roles using (user_id)").
		Join("roles using (role_id)").
		Where(sq.Eq{"user_name": username})

	sql, args, err := query.ToSql()
	logger.Logger.Sugar().Infow(userTable, sql, args)

	if err != nil {
		err = errors.NewInternal(err, "create query")
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
			&user.CreateAt,
			&user.Role)

	if err != nil {
		logger.Error(err.Error())
		if !errors.Is(err, pgx.ErrNoRows) {
			err = errors.NewInternal(err, "executy query")
			return model.User{}, err
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *repo) GetByID(ctx context.Context, ID int) (model.User, error) {
	ctx, span := tracer.Start(ctx, "repo.GetByID")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	query := r.queryBuilder.
		Select(
			"user_id",
			"user_name",
			"user_email",
			"user_password",
			"create_at",
			"role_name").
		From(userTable).
		Join("user_roles using (user_id)").
		Join("roles using (role_id)").
		Where(sq.Eq{"user_id": ID})

	sql, args, err := query.ToSql()
	logger.Logger.Sugar().Infow(userTable, sql, args)

	if err != nil {
		err = errors.NewInternal(err, "create query")
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
			&user.CreateAt,
			&user.Role)
	if err != nil {
		logger.Error(err.Error())
		if !errors.Is(err, pgx.ErrNoRows) {
			err = errors.NewInternal(err, "executy query")
			return model.User{}, err
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *repo) Delete(ctx context.Context, username string) error {
	ctx, span := tracer.Start(ctx, "repo.Delete")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	sql, args, err := r.queryBuilder.
		Delete(userTable).
		Where(sq.Eq{"user_name": username}).
		ToSql()

	logger.Logger.Sugar().Infow(userTable, sql, args)

	if err != nil {
		err = errors.NewInternal(err, "create query")
		logger.Error(err.Error())
		return err
	}

	if exec, execErr := r.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = errors.NewInternal(execErr, "execute query")
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Delete() {
		logger.Error(execErr.Error())
		return execErr
	}

	return nil
}

func (r *repo) Update(ctx context.Context, user model.User) error {
	ctx, span := tracer.Start(ctx, "repo.Update")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	infoStorageMap := toUpdateStorageMap(user)

	sql, args, buildErr := r.queryBuilder.
		Update(userTable).
		SetMap(infoStorageMap).
		Where(sq.Eq{"user_name": user.Username}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	logger.Logger.Sugar().Infow(userTable, sql, args)

	if buildErr != nil {
		buildErr = errors.NewInternal(buildErr, "create query")
		logger.Error(buildErr.Error())
		return buildErr
	}

	if exec, execErr := r.client.Exec(ctx, sql, args...); execErr != nil {
		execErr = errors.NewInternal(execErr, "execute query")
		logger.Error(execErr.Error())
		return execErr
	} else if exec.RowsAffected() == 0 || !exec.Update() {
		logger.Error(execErr.Error())
		return execErr
	}

	return nil
}
