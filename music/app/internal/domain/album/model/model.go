package model

import (
	"errors"
	"time"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	creditModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	infoModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	trackModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/google/uuid"
)

var (
	ErrValidation = errors.New("title must not be empty")
)

type Album struct {
	ID         string
	Title      string
	ReleasedAt int64
	CreatedAt  int64
}

type FullAlbum struct {
	Album     Album
	Info      infoModel.Info
	Tracklist []trackModel.Track
	Credits   []creditModel.Credit
}

func (a Album) ToProto() *albumPb.Album {
	return &albumPb.Album{
		AlbumId:    a.ID,
		Title:      a.Title,
		CreatedAt:  a.CreatedAt,
		ReleasedAt: a.ReleasedAt,
	}
}

func NewAlbumFromPB(pb *albumPb.CreateAlbumRequest) *FullAlbum {
	protoList := pb.GetTracklist()
	tracklist := make([]trackModel.Track, len(protoList))

	protoCredits := pb.GetCredits()
	credits := make([]creditModel.Credit, len(protoCredits))

	for i := 0; i < len(protoList); i++ {
		tracklist[i] = trackModel.NewTrackFromPB(protoList[i])
	}

	for i := 0; i < len(protoCredits); i++ {
		credits[i] = creditModel.NewCreditFromPB(protoCredits[i])
	}

	return &FullAlbum{
		Album: Album{
			ID:         uuid.New().String(),
			Title:      pb.GetTitle(),
			ReleasedAt: pb.GetReleasedAt(),
			CreatedAt:  time.Now().UnixMilli(),
		},
		Info:      infoModel.NewInfoFromPB(pb),
		Tracklist: tracklist,
		Credits:   credits,
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
