package model

import (
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/mitchellh/mapstructure"
)

type Profession struct {
	ID    int64
	Title string
}

func ToMap(p *Profession) (map[string]interface{}, error) {
	var updateProfMap map[string]interface{}
	err := mapstructure.Decode(p, &updateProfMap)
	if err != nil {
		return updateProfMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updateProfMap, nil
}
