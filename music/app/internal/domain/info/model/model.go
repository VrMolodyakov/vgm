package model

import (
	infoPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/info/v1"
	"github.com/google/uuid"
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

func NewInfoFromPB(pb *infoPb.CreateAlbumInfoRequest) Info {
	return Info{
		ID:             uuid.New().String(),
		AlbumID:        pb.GetAlbumId(),
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
