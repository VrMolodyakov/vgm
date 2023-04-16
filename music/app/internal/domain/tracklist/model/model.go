package model

import (
	"errors"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
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

func (t *Track) IsValid() bool {
	return t.Title != ""
}

func (t *Track) ToProto() albumPb.TrackInfo {
	return albumPb.TrackInfo{
		Id:       t.ID,
		AlbumId:  t.AlbumID,
		Title:    t.Title,
		Duration: t.Duration,
	}
}

func NewTrackFromPB(pb *albumPb.Track) Track {
	return Track{
		Title:    pb.Title,
		Duration: pb.Duration,
	}
}
