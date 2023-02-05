package tracklist

// import (
// 	"context"

// 	trackPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/track/v1"
// 	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
// )

// func (s *server) CreateTracklist(ctx context.Context, request *trackPb.CreateTracklistRequest) (*trackPb.CreateTracklistResponse, error) {
// 	tracklistPb := request.GetTrack()
// 	tracklist := make([]model.Track, len(tracklistPb))
// 	for i := 0; i < len(tracklistPb); i++ {
// 		tracklist[i] = model.NewTrackFromPB(tracklistPb[i])
// 	}
// 	err := s.tracklistPolicy.Create(ctx, tracklist)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &trackPb.CreateTracklistResponse{}, nil
// }
