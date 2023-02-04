package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
)

type TrackDAO interface {
	Create(ctx context.Context, tracklist []model.Track) error
	GetOne(ctx context.Context, albumID string) (dao.TrackStorage, error)
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

// func (s *trackService) GetOne(ctx context.Context, infoID string) (model.Track, error) {
// 	if infoID == "" TrackDAO
// 		err := errors.New("id must not be empty")
// 		logging.LoggerFromContext(ctx).Error(err.Error())
// 		return model.Track{}, err
// 	}
// 	dbInfo, err := s.infoDAO.GetOne(ctx, infoID)
// 	if err != nil {
// 		return model.Track{}, errors.Wrap(err, "trackService.Create")
// 	}
// TrackDAO dbInfo.ToModel(), nil
// }
