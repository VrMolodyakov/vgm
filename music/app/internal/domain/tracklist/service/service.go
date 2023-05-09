package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("track-service")
)

type TrackRepo interface {
	Create(ctx context.Context, tracklist []model.Track) error
	GetAll(ctx context.Context, albumID string) ([]model.Track, error)
	Update(ctx context.Context, track model.Track) error
	Delete(ctx context.Context, id string) error
}

type trackService struct {
	trackRepo TrackRepo
}

func NewTrackService(dao TrackRepo) *trackService {
	return &trackService{trackRepo: dao}
}

func (t *trackService) Create(ctx context.Context, tracklist []model.Track) error {
	ctx, span := tracer.Start(ctx, "service.Create")
	defer span.End()

	for _, t := range tracklist {
		if !t.IsValid() {
			return model.ErrValidation
		}
	}

	err := t.trackRepo.Create(ctx, tracklist)
	if err != nil {
		return errors.Wrap(err, "trackService.Create")
	}
	return nil

}

func (t *trackService) GetAll(ctx context.Context, albumID string) ([]model.Track, error) {
	ctx, span := tracer.Start(ctx, "service.GetAll")
	defer span.End()

	if albumID == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return nil, err
	}
	trakclist, err := t.trackRepo.GetAll(ctx, albumID)
	if err != nil {
		return nil, errors.Wrap(err, "trackService.Create")
	}
	return trakclist, nil
}

func (t *trackService) Update(ctx context.Context, track model.Track) error {
	ctx, span := tracer.Start(ctx, "service.Update")
	defer span.End()

	if !track.IsValid() {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err
	}
	err := t.trackRepo.Update(ctx, track)
	if err != nil {
		return errors.Wrap(err, "trackService.Update")
	}
	return nil
}

func (t *trackService) Delete(ctx context.Context, id string) error {
	ctx, span := tracer.Start(ctx, "service.Delete")
	defer span.End()

	if id == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err

	}
	return t.trackRepo.Delete(ctx, id)
}
