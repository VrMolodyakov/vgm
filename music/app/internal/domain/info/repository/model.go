package reposotory

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	mapper "github.com/worldline-go/struct2"
)

const (
	fields = 10
)

type InfoStorage struct {
	ID             string  `struct:"album_info_id"`
	AlbumID        string  `struct:"album_id"`
	CatalogNumber  string  `struct:"catalog_number" `
	ImageSrc       string  `struct:"image_srs" `
	Barcode        string  `struct:"barcode" `
	CurrencyCode   string  `struct:"currency_code" `
	MediaFormat    string  `struct:"media_format" `
	Classification string  `struct:"classification"`
	Publisher      string  `struct:"publisher"`
	Price          float64 `struct:"price" `
}

func toStorageMap(info *model.Info) map[string]interface{} {
	storage := fromModel(info)
	infoStorageMap := (&mapper.Decoder{}).Map(storage)
	return infoStorageMap
}

func fromModel(m *model.Info) InfoStorage {
	return InfoStorage{
		ID:             m.ID,
		AlbumID:        m.AlbumID,
		CatalogNumber:  m.CatalogNumber,
		ImageSrc:       m.ImageSrc,
		Barcode:        m.Barcode,
		CurrencyCode:   m.CurrencyCode,
		MediaFormat:    m.MediaFormat,
		Classification: m.Classification,
		Publisher:      m.Publisher,
		Price:          m.Price,
	}
}

func (i *InfoStorage) toModel() model.Info {
	return model.Info{
		ID:             i.ID,
		AlbumID:        i.AlbumID,
		CatalogNumber:  i.CatalogNumber,
		ImageSrc:       i.ImageSrc,
		Barcode:        i.Barcode,
		CurrencyCode:   i.CurrencyCode,
		MediaFormat:    i.MediaFormat,
		Classification: i.Classification,
		Publisher:      i.Publisher,
		Price:          i.Price,
	}
}

func toUpdateStorageMap(m *model.Info) map[string]interface{} {

	storageMap := make(map[string]interface{}, fields)

	if m.CatalogNumber != "" {
		storageMap["catalog_number"] = m.CatalogNumber
	}
	if m.ImageSrc != "" {
		storageMap["image_srs"] = m.ImageSrc
	}
	if m.Barcode != "" {
		storageMap["barcode"] = m.Barcode
	}
	if m.CurrencyCode != "" {
		storageMap["currency_code"] = m.CurrencyCode
	}
	if m.MediaFormat != "" {
		storageMap["media_format"] = m.MediaFormat
	}
	if m.Classification != "" {
		storageMap["classification"] = m.Classification
	}
	if m.Publisher != "" {
		storageMap["publisher"] = m.Publisher
	}
	if m.Price != 0 {
		storageMap["price"] = m.Price
	}

	return storageMap
}
