package model

import (
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/mitchellh/mapstructure"
)

type Credit struct {
	PersonID     string `mapstructure:"album_id"`
	AlbumID      string `mapstructure:"album_id"`
	ProfessionID string `mapstructure:"album_id"`
}

func ToMap(a *Credit) (map[string]interface{}, error) {
	var updateAlbumMap map[string]interface{}
	err := mapstructure.Decode(a, &updateAlbumMap)
	if err != nil {
		return updateAlbumMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updateAlbumMap, nil
}
