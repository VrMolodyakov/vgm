package model

import (
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/mitchellh/mapstructure"
)

type Tracklist struct {
	ID      int64  `mapstructure:"track_id "`
	AlbumID string `mapstructure:"album_id"`
	Title   string `mapstructure:"title"`
}

func ToMap(t *Tracklist) (map[string]interface{}, error) {
	var updateTrackMap map[string]interface{}
	err := mapstructure.Decode(t, &updateTrackMap)
	if err != nil {
		return updateTrackMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updateTrackMap, nil
}
