package model

import (
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/mitchellh/mapstructure"
)

type Credit struct {
	PersonID     int64  `mapstructure:"person_id"`
	AlbumID      string `mapstructure:"album_id"`
	ProfessionID int64  `mapstructure:"profession_id"`
}

func ToMap(a *Credit) (map[string]interface{}, error) {
	var updateCreditMap map[string]interface{}
	err := mapstructure.Decode(a, &updateCreditMap)
	if err != nil {
		return updateCreditMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updateCreditMap, nil
}
