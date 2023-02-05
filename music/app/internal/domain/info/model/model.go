package model

import (
	"errors"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	infoPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/info/v1"
	"github.com/google/uuid"
)

var (
	ErrValidation = errors.New("album id must not be empty")
)

type Info struct {
	ID             string
	AlbumID        string
	CatalogNumber  string
	ImageSrc       string
	Barcode        string
	CurrencyCode   string
	MediaFormat    string
	Classification string
	Publisher      string
	Price          float64
}

// func NewInfoFromPB(pb *infoPb.CreateAlbumInfoRequest) Info {
// 	return Info{
// 		ID:             uuid.New().String(),
// 		AlbumID:        pb.GetAlbumId(),
// 		CatalogNumber:  pb.GetCatalogNumber(),
// 		ImageSrc:       pb.GetImageSrc(),
// 		Barcode:        pb.GetBarcode(),
// 		CurrencyCode:   pb.GetCurrencyCode(),
// 		MediaFormat:    pb.GetMediaFormat(),
// 		Classification: pb.GetClassification(),
// 		Publisher:      pb.GetPublisher(),
// 		Price:          pb.GetPrice(),
// 	}
// }

func NewInfoFromPB(pb *albumPb.CreateFullAlbumRequest) Info {
	return Info{
		ID:             uuid.New().String(),
		CatalogNumber:  pb.GetCatalogNumber(),
		ImageSrc:       pb.GetImageSrc(),
		Barcode:        pb.GetBarcode(),
		CurrencyCode:   pb.GetCurrencyCode(),
		MediaFormat:    pb.GetMediaFormat(),
		Classification: pb.GetClassification(),
		Publisher:      pb.GetPublisher(),
		Price:          pb.GetPrice(),
	}
}

func (i *Info) ToProto() *infoPb.Info {
	return &infoPb.Info{
		AlbumInfoId:    i.ID,
		AlbumId:        i.AlbumID,
		CatalogNumber:  i.CatalogNumber,
		ImageSrc:       &i.ImageSrc,
		Barcode:        &i.Barcode,
		Price:          i.Price,
		CurrencyCode:   i.CurrencyCode,
		MediaFormat:    i.MediaFormat,
		Classification: i.Classification,
		Publisher:      i.Publisher,
	}
}

func UpdateModelFromPB(pb *infoPb.UpdateAlbumInfoRequest) Info {
	var info Info

	info.ID = pb.GetAlbumInfoId()

	if pb.CatalogNumber != nil {
		info.CatalogNumber = pb.GetCatalogNumber()
	}
	if pb.ImageSrc != nil {
		info.ImageSrc = pb.GetImageSrc()
	}
	if pb.Barcode != nil {
		info.Barcode = pb.GetBarcode()
	}
	if pb.Price != nil {
		info.Price = pb.GetPrice()
	}
	if pb.CurrencyCode != nil {
		info.CurrencyCode = pb.GetCurrencyCode()
	}
	if pb.MediaFormat != nil {
		info.MediaFormat = pb.GetMediaFormat()
	}
	if pb.Classification != nil {
		info.Classification = pb.GetClassification()
	}
	if pb.Publisher != nil {
		info.Publisher = pb.GetPublisher()
	}

	return info
}

func (i *Info) IsEmpty() bool {
	return i.ID == ""
}
