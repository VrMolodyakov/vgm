package reposotory

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type tracklistDAO struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "track"
)

func NewTracklistStorage(client db.PostgreSQLClient) *tracklistDAO {
	return &tracklistDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (t *tracklistDAO) Create(ctx context.Context, tracklist []model.Track) error {
	logger := logging.LoggerFromContext(ctx)
	insertState := t.queryBuilder.Insert(table).Columns("album_id", "title", "duration")
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

func (t *tracklistDAO) GetAll(ctx context.Context, albumID string) ([]model.Track, error) {
	logger := logging.LoggerFromContext(ctx)

	query := t.queryBuilder.
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

	rows, queryErr := t.client.Query(ctx, sql, args...)
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
