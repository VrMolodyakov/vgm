package dao

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	mapper "github.com/worldline-go/struct2"
)

type InfoStorage struct {
	ID             string  `struct:"album_info_id"`
	AlbumID        string  `struct:"album_id"`
	CatalogNumber  string  `struct:"catalog_number" `
	ImageSrc       string  `struct:"image_srs" `
	Barcode        string  `struct:"barcode" `
	CurrencyCode   string  `struct:"currency_id" `
	MediaFormat    string  `struct:"media_format" `
	Classification string  `struct:"classification"`
	Publisher      string  `struct:"publisher"`
	Price          float64 `struct:"price" `
}

func ToStorageMap(info model.Info) map[string]interface{} {
	infoStorageMap := (&mapper.Decoder{}).Map(info)
	return infoStorageMap
}
