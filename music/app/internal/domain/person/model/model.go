package model

import (
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	ID    int64  `mapstructure:"profession_id"`
	Title string `mapstructure:"professional_title"`
}

func ToMap(p *Person) (map[string]interface{}, error) {
	var updatePersonMap map[string]interface{}
	err := mapstructure.Decode(p, &updatePersonMap)
	if err != nil {
		return updatePersonMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updatePersonMap, nil
}
