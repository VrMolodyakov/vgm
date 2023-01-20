package model

import (
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/mitchellh/mapstructure"
)

type Currency struct {
	ID     int64  `mapstructure:"currency_id"`
	Name   string `mapstructure:"name"`
	Symbol string `mapstructure:"symbol"`
}

func ToMap(c *Currency) (map[string]interface{}, error) {
	var updateCurrencyMap map[string]interface{}
	err := mapstructure.Decode(c, &updateCurrencyMap)
	if err != nil {
		return updateCurrencyMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updateCurrencyMap, nil
}
