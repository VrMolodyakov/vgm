package tracklist

import (
	"context"

	trackPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/track/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
)

type TracklistPolicy interface {
	Create(ctx context.Context, tracklist []model.Track) error
	GetOne(ctx context.Context, trackID string) (model.Track, error)
}

type server struct {
	tracklistPolicy TracklistPolicy
	trackPb.UnimplementedTrackServiceServer
}

func NewServer(policy TracklistPolicy, s trackPb.UnimplementedTrackServiceServer) *server {
	return &server{
		tracklistPolicy:                 policy,
		UnimplementedTrackServiceServer: s,
	}
}
