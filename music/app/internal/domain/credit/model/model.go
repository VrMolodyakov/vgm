package model

import (
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/mitchellh/mapstructure"
)

type Credit struct {
	PersonID     int64
	AlbumID      string
	ProfessionID int64
}

func ToMap(a *Credit) (map[string]interface{}, error) {
	var updateCreditMap map[string]interface{}
	err := mapstructure.Decode(a, &updateCreditMap)
	if err != nil {
		return updateCreditMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updateCreditMap, nil
}

// `struct:"person_id"`
// `struct:"album_id"`
// `struct:"profession_id"`
