package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
)

type TrackService interface {
	Create(ctx context.Context, tracklist []model.Track) error
	GetAll(ctx context.Context, albumID string) ([]model.Track, error)
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
