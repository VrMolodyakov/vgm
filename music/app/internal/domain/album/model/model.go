package model

import (
	"errors"
	"time"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"github.com/google/uuid"
)

var (
	ErrValidation = errors.New("Title must not be empty")
)

type Album struct {
	ID         string
	Title      string
	ReleasedAt int64
	CreatedAt  int64
}

func (a Album) ToProto() *albumPb.Album {
	return &albumPb.Album{
		AlbumId:    a.ID,
		Title:      a.Title,
		CreatedAt:  a.CreatedAt,
		ReleasedAt: a.ReleasedAt,
	}
}

func NewAlbumFromPB(pb *albumPb.CreateFullAlbumRequest) Album {
	return Album{
		ID:         uuid.New().String(),
		Title:      pb.GetTitle(),
		ReleasedAt: pb.GetReleasedAt(),
		CreatedAt:  time.Now().UnixMilli(),
	}
}

func UpdateModelFromPB(pb *albumPb.UpdateAlbumRequest) Album {
	var album Album

	if pb.Title != nil {
		album.Title = pb.GetTitle()
	}

	if pb.CreatedAt != nil {
		album.CreatedAt = pb.GetCreatedAt()
	}

	if pb.ReleasedAt != nil {
		album.ReleasedAt = pb.GetReleasedAt()
	}

	album.ID = pb.GetId()
	return album
}

func (a *Album) IsEmpty() bool {
	return a.Title == ""
}
