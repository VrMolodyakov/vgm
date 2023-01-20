package model

import (
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/mitchellh/mapstructure"
)

type Info struct {
	ID             string  `mapstructure:"album_info_id"`
	CatalogNumber  string  `mapstructure:"catalog_number" `
	ImageSrc       string  `mapstructure:"image_srs" `
	Barcode        string  `mapstructure:"barcode" `
	ReleaseDate    int64   `mapstructure:"release_date"`
	Price          float64 `mapstructure:"price" `
	CurrencyID     int     `mapstructure:"currency_id" `
	MediaFormat    string  `mapstructure:"media_format" `
	Classification string  `mapstructure:"classification"`
	Publisher      string  `mapstructure:"publisher"`
}

func ToMap(c *Info) (map[string]interface{}, error) {
	var updateInfoMap map[string]interface{}
	err := mapstructure.Decode(c, &updateInfoMap)
	if err != nil {
		return updateInfoMap, errors.Wrap(err, "mapstructure.Decode(product)")
	}

	return updateInfoMap, nil
}
