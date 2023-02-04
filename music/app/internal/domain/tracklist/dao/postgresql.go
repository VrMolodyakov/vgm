package dao

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
	table = "tracklist"
)

func NewTracklistStorage(client db.PostgreSQLClient) *tracklistDAO {
	return &tracklistDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (t *tracklistDAO) Create(ctx context.Context, tracklist []model.Track) error {
	logger := logging.LoggerFromContext(ctx)
	insertState := t.queryBuilder.Insert(table).Columns("album_id", "title")
	for _, track := range tracklist {
		insertState = insertState.Values(track.AlbumID, track.Title)
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

func (t *tracklistDAO) GetOne(ctx context.Context, albumID string) (TrackStorage, error) {
	logger := logging.LoggerFromContext(ctx)

	query := t.queryBuilder.
		Select(
			"track_id",
			"album_id",
			"title").
		From(table).
		Where(sq.Eq{"album_id": albumID})

	sql, args, err := query.ToSql()

	logger.Infow(table, sql, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err.Error())
		return TrackStorage{}, err
	}

	var tS TrackStorage
	err = t.client.QueryRow(ctx, sql, args...).
		Scan(
			&tS.ID,
			&tS.AlbumID,
			&tS.Title)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err.Error())
		return TrackStorage{}, err
	}
	return tS, nil
}
