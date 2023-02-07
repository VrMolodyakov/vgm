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
	GetAll(ctx context.Context, albumID string) ([]dao.TrackStorage, error)
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

func (t *trackService) GetAll(ctx context.Context, albumID string) ([]model.Track, error) {
	if albumID == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return nil, err
	}
	storageList, err := t.trackDAO.GetAll(ctx, albumID)
	if err != nil {
		return nil, errors.Wrap(err, "trackService.Create")
	}
	trakclist := make([]model.Track, len(storageList))
	for i := 0; i < len(storageList); i++ {
		trakclist[i] = storageList[i].ToModel()
	}
	return trakclist, nil
}
