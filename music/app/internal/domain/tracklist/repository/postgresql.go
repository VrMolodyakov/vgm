package reposotory

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"go.opentelemetry.io/otel"
)

type repo struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "track"
)

var (
	tracer = otel.Tracer("track-repo")
)

func NewTracklistRepo(client db.PostgreSQLClient) *repo {
	return &repo{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (r *repo) Create(ctx context.Context, tracklist []model.Track) error {
	_, span := tracer.Start(ctx, "repo.Create")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	insertState := r.queryBuilder.Insert(table).Columns("album_id", "title", "duration")
	for _, track := range tracklist {
		insertState = insertState.Values(track.AlbumID, track.Title, track.Duration)
	}
	sql, args, err := insertState.ToSql()
	logger.Infow(table, sql, args)
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
		execErr = db.ErrDoQuery(errors.New("tracklist was not created. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil
}

func (r *repo) GetAll(ctx context.Context, albumID string) ([]model.Track, error) {
	_, span := tracer.Start(ctx, "repo.GetAll")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)

	query := r.queryBuilder.
		Select(
			"track_id",
			"album_id",
			"title",
			"duration").
		From(table).
		Where(sq.Eq{"album_id": albumID})

	sql, args, err := query.ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return nil, err
	}

	rows, queryErr := r.client.Query(ctx, sql, args...)
	if queryErr != nil {
		queryErr = db.ErrDoQuery(queryErr)
		logger.Error(queryErr.Error())
		return nil, queryErr
	}
	tracklist := make([]model.Track, 0)
	for rows.Next() {
		storage := TrackStorage{}
		if queryErr = rows.Scan(
			&storage.ID,
			&storage.AlbumID,
			&storage.Title,
			&storage.Duration,
		); queryErr != nil {
			queryErr = db.ErrScan(queryErr)
			logger.Error(queryErr.Error())
			return nil, queryErr
		}
		tracklist = append(tracklist, storage.toModel())

	}
	return tracklist, nil
}

func (r *repo) Update(ctx context.Context, track model.Track) error {
	_, span := tracer.Start(ctx, "repo.Update")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	infoStorageMap := toUpdateStorageMap(&track)

	sql, args, buildErr := r.queryBuilder.
		Update(table).
		SetMap(infoStorageMap).
		Where(sq.Eq{"album_info_id": track.ID}).
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
		execErr = db.ErrDoQuery(errors.New("track was not updated. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	_, span := tracer.Start(ctx, "repo.Delete")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	sql, args, buildErr := r.queryBuilder.
		Delete(table).
		Where(sq.Eq{"album_id": id}).
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
		execErr = db.ErrDoQuery(errors.New("track was not deleted. 0 rows were affected"))
		logger.Error(execErr.Error())
		return execErr
	}

	return nil

}
