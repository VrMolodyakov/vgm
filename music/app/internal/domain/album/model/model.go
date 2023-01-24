package model

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/dao"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/mitchellh/mapstructure"
)

type Album struct {
	ID       string `mapstructure:"album_id"`
	Title    string `mapstructure:"title"`
	CreateAt int64  `mapstructure:"create_at"`
}

func ToMap(a *Album) (map[string]interface{}, error) {
	var updateAlbumMap map[string]interface{}
	err := mapstructure.Decode(a, &updateAlbumMap)
	if err != nil {
		return updateAlbumMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updateAlbumMap, nil
}

func NewAlbum(albumDao dao.AlbumDAO) Album {
	return Album{}
}
