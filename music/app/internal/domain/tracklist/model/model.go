package model

import (
	"errors"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
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

func NewTrackFromPB(pb *albumPb.Track) Track {
	return Track{
		Title:    pb.Title,
		Duration: pb.Duration,
	}
}

func (t *Track) ToProto() *trackPb.Track {
	return &trackPb.Track{
		Title:    t.Title,
		Duration: t.Duration,
		AlbumId:  t.AlbumID,
	}
}
