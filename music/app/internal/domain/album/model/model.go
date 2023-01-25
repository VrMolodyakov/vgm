package model

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/dao"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
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

func NewAlbum(album dao.AlbumStorage) Album {
	return Album{
		ID:         album.ID,
		Title:      album.Title,
		CreatedAt:  album.CreatedAt.Unix(),
		ReleasedAt: album.ReleasedAt.Unix(),
	}
}
