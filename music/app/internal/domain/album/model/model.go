package model

import (
	"errors"
	"time"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	creditModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	trackModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/google/uuid"
)

var (
	ErrValidation = errors.New("title must not be empty")
)

type AlbumView struct {
	ID         string
	Title      string
	ReleasedAt int64
	CreatedAt  int64
}

type AlbumInfo struct {
	Album AlbumView
	Info  Info
}

type Album struct {
	Album     AlbumView
	Info      Info
	Tracklist []trackModel.Track
	Credits   []creditModel.Credit
}

type FullAlbum struct {
	Album     AlbumView
	Info      Info
	Tracklist []trackModel.Track
	Credits   []creditModel.CreditInfo
}

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

func (a AlbumView) ToProto() *albumPb.Album {
	return &albumPb.Album{
		AlbumId:    a.ID,
		Title:      a.Title,
		CreatedAt:  a.CreatedAt,
		ReleasedAt: a.ReleasedAt,
	}
}

func NewAlbumFromPB(pb *albumPb.CreateAlbumRequest) *Album {
	protoList := pb.GetTracklist()
	tracklist := make([]trackModel.Track, len(protoList))

	protoCredits := pb.GetCredits()
	credits := make([]creditModel.Credit, len(protoCredits))

	albumView := AlbumView{
		ID:         uuid.New().String(),
		Title:      pb.GetTitle(),
		ReleasedAt: pb.GetReleasedAt(),
		CreatedAt:  time.Now().UnixMilli(),
	}
	info := NewInfoFromPB(pb)
	info.AlbumID = albumView.ID

	for i := 0; i < len(protoList); i++ {
		tracklist[i] = trackModel.NewTrackFromPB(protoList[i])
		tracklist[i].AlbumID = albumView.ID
	}

	for i := 0; i < len(protoCredits); i++ {
		credits[i] = creditModel.NewCreditFromPB(protoCredits[i])
		credits[i].AlbumID = albumView.ID
	}

	return &Album{
		Album:     albumView,
		Info:      info,
		Tracklist: tracklist,
		Credits:   credits,
	}
}

func NewInfoFromPB(pb *albumPb.CreateAlbumRequest) Info {
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

func UpdateModelFromPB(pb *albumPb.UpdateAlbumRequest) AlbumView {
	var album AlbumView

	if pb.Title != nil {
		album.Title = pb.GetTitle()
	}

	if pb.CreatedAt != nil {
		album.CreatedAt = pb.GetCreatedAt()
	}

	if pb.ReleasedAt != nil {
		album.ReleasedAt = pb.GetReleasedAt()
	}

	album.ID = pb.GetId()
	return album
}

func (a *AlbumView) IsEmpty() bool {
	return a.Title == ""
}

func (a *Album) IsEmpty() bool {
	return a.Album.Title == ""
}
