package reposotory

import (
	"time"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	mapper "github.com/worldline-go/struct2"
)

const (
	fields = 3
)

type AlbumViewStorage struct {
	ID         string    `struct:"album_id"`
	Title      string    `struct:"title"`
	ReleasedAt time.Time `struct:"released_at"`
	CreatedAt  time.Time `struct:"created_at"`
}

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

type AlbumInfoStorage struct {
	Album AlbumViewStorage
	Info  InfoStorage
}

func ToInfoStorageMap(info *model.Info) map[string]interface{} {
	storage := fromInfoModel(info)
	infoStorageMap := (&mapper.Decoder{}).Map(storage)
	return infoStorageMap
}

func fromInfoModel(info *model.Info) InfoStorage {
	return InfoStorage{
		ID:             info.ID,
		AlbumID:        info.AlbumID,
		CatalogNumber:  info.CatalogNumber,
		ImageSrc:       info.ImageSrc,
		Barcode:        info.Barcode,
		CurrencyCode:   info.CurrencyCode,
		MediaFormat:    info.MediaFormat,
		Classification: info.Classification,
		Publisher:      info.Publisher,
		Price:          info.Price,
	}
}

func fromModel(album model.AlbumView) AlbumViewStorage {
	createdAt := time.UnixMilli(album.CreatedAt)
	releasedAt := time.UnixMilli(album.ReleasedAt)
	return AlbumViewStorage{
		ID:         album.ID,
		Title:      album.Title,
		CreatedAt:  createdAt,
		ReleasedAt: releasedAt,
	}
}

func (album AlbumViewStorage) toModel() model.AlbumView {
	return model.AlbumView{
		ID:         album.ID,
		Title:      album.Title,
		CreatedAt:  album.CreatedAt.UnixMilli(),
		ReleasedAt: album.ReleasedAt.UnixMilli(),
	}
}

func (album *AlbumInfoStorage) toModel() model.AlbumInfo {
	return model.AlbumInfo{
		Album: model.AlbumView{
			ID:         album.Album.ID,
			Title:      album.Album.Title,
			CreatedAt:  album.Album.CreatedAt.UnixMilli(),
			ReleasedAt: album.Album.ReleasedAt.UnixMilli(),
		},
		Info: model.Info{
			ID:             album.Info.ID,
			AlbumID:        album.Info.AlbumID,
			CatalogNumber:  album.Info.CatalogNumber,
			ImageSrc:       album.Info.ImageSrc,
			Barcode:        album.Info.Barcode,
			CurrencyCode:   album.Info.CurrencyCode,
			MediaFormat:    album.Info.MediaFormat,
			Classification: album.Info.Classification,
			Publisher:      album.Info.Publisher,
			Price:          album.Info.Price,
		},
	}
}

func ToStorageMap(album model.AlbumView) map[string]interface{} {
	storage := fromModel(album)
	albumStorageMap := (&mapper.Decoder{}).Map(storage)
	return albumStorageMap
}

func ToUpdateStorageMap(album model.AlbumView) map[string]interface{} {
	m := make(map[string]interface{}, fields)

	if album.Title != "" {
		m["title"] = album.Title
	}

	if album.CreatedAt != 0 {
		m["created_at"] = time.UnixMilli(album.CreatedAt)
	}

	if album.ReleasedAt != 0 {
		m["released_at"] = time.UnixMilli(album.ReleasedAt)
	}

	return m
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
