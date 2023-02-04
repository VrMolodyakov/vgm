package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type TrackDAO interface {
	Create(ctx context.Context, tracklist []model.Track) error
	GetOne(ctx context.Context, trackID string) (dao.TrackStorage, error)
}

type trackService struct {
	trackDAO TrackDAO
}

func NewTrackService(dao TrackDAO) *trackService {
	return &trackService{trackDAO: dao}
}

func (t *trackService) Create(ctx context.Context, tracklist []model.Track) error {
	for _, t := range tracklist {
		if t.IsEmpty() {
			return model.ErrValidation
		}
	}

	err := t.trackDAO.Create(ctx, tracklist)
	if err != nil {
		return errors.Wrap(err, "trackService.Create")
	}
	return nil

}

func (t *trackService) GetOne(ctx context.Context, trackID string) (model.Track, error) {
	if trackID == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return model.Track{}, err
	}
	track, err := t.trackDAO.GetOne(ctx, trackID)
	if err != nil {
		return model.Track{}, errors.Wrap(err, "trackService.Create")
	}
	return track.ToModel(), nil
}
