package model

import (
	"errors"

	trackPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/track/v1"
)

var (
	ErrValidation = errors.New("title must not be empty")
)

type Track struct {
	ID       int64
	AlbumID  string
	Title    string
	Duration string
}

func (t *Track) IsEmpty() bool {
	return t.Title == ""
}

func NewTrackFromPB(pb *trackPb.Track) Track {
	return Track{
		AlbumID:  pb.AlbumId,
		Title:    pb.Title,
		Duration: pb.Duration,
	}
}
