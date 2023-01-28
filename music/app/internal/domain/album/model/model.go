package model

import (
	"fmt"
	"time"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"github.com/google/uuid"
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

func NewAlbumFromPB(pb *albumPb.CreateAlbumRequest) Album {
	unix := time.Now().Unix()
	fmt.Println(time.Unix(unix, 0))
	return Album{
		ID:         uuid.New().String(),
		Title:      pb.GetTitle(),
		ReleasedAt: pb.GetReleaseAt(),
		CreatedAt:  time.Now().UnixMilli(),
	}
}
