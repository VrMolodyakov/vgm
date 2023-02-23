package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
)

type TrackService interface {
	Create(ctx context.Context, tracklist []model.Track) error
	GetAll(ctx context.Context, albumID string) ([]model.Track, error)
	Update(ctx context.Context, track model.Track) error
	Delete(ctx context.Context, id string) error
}

type trackPolicy struct {
	trackService TrackService
}

func NewTrackPolicy(service TrackService) *trackPolicy {
	return &trackPolicy{trackService: service}
}

func (p *trackPolicy) Create(ctx context.Context, tracklist []model.Track) error {
	return p.trackService.Create(ctx, tracklist)
}

func (p *trackPolicy) GetAll(ctx context.Context, albumID string) ([]model.Track, error) {
	return p.trackService.GetAll(ctx, albumID)
}

func (p *trackPolicy) Update(ctx context.Context, track model.Track) error {
	return p.trackService.Update(ctx, track)
}

func (p *trackPolicy) Delete(ctx context.Context, albumID string) error {
	return p.trackService.Delete(ctx, albumID)
}
