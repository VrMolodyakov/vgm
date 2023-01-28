package model

import (
	"fmt"
	"time"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type Album struct {
	ID         string `mapstructure:"album_id"`
	Title      string `mapstructure:"title"`
	ReleasedAt int64  `mapstructure:"released_at"`
	CreatedAt  int64  `mapstructure:"created_at"`
}

func (a Album) ToMap() (map[string]interface{}, error) {
	var updateAlbumMap map[string]interface{}
	err := mapstructure.Decode(a, &updateAlbumMap)
	if err != nil {
		return updateAlbumMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updateAlbumMap, nil
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
