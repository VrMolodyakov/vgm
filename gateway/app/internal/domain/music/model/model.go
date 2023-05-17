package model

import (
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/music/dto"
	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
)

type Album struct {
	Album     AlbumView
	Info      Info
	Tracklist []Track
	Credits   []Credit
}

type AlbumView struct {
	ID         string
	Title      string
	ReleasedAt int64
	CreatedAt  int64
}

type AlbumPreview struct {
	ID            string
	Title         string
	ReleasedAt    int64
	CreatedAt     int64
	SmallImageSrc string
	Publisher     string
}

type Person struct {
	FirstName string
	LastName  string
	BirthDate int64
}

type Info struct {
	ID             string
	AlbumID        string
	CatalogNumber  string
	FullImageSrc   string
	SmallImageSrc  string
	Barcode        string
	CurrencyCode   string
	MediaFormat    string
	Classification string
	Publisher      string
	Price          float64
}

type Track struct {
	ID       int64
	AlbumID  string
	Title    string
	Duration string
}

type CreditInfo struct {
	FirstName  string
	LastName   string
	Profession string
}

type Credit struct {
	PersonID   int64
	AlbumID    string
	Profession string
}

type FullAlbum struct {
	Album     AlbumView
	Info      Info
	Tracklist []Track
	Credits   []CreditInfo
}

func AlbumFromDto(dto dto.AlbumReq) Album {
	tracklist := make([]Track, len(dto.Tracklist))
	credits := make([]Credit, len(dto.Credits))

	for i := 0; i < len(tracklist); i++ {
		tracklist[i] = Track{
			Title:    dto.Tracklist[i].Title,
			Duration: dto.Tracklist[i].Duration,
		}
	}

	for i := 0; i < len(credits); i++ {
		credits[i] = Credit{
			PersonID:   dto.Credits[i].PersonID,
			Profession: dto.Credits[i].Profession,
		}
	}

	return Album{
		Album: AlbumView{
			Title:      dto.Album.Title,
			ReleasedAt: dto.Album.ReleasedAt,
			CreatedAt:  dto.Album.CreatedAt,
		},
		Info: Info{
			CatalogNumber:  dto.Info.CatalogNumber,
			FullImageSrc:   dto.Info.FullImageSrc,
			SmallImageSrc:  dto.Info.SmallImageSrc,
			Barcode:        dto.Info.Barcode,
			CurrencyCode:   dto.Info.CurrencyCode,
			MediaFormat:    dto.Info.MediaFormat,
			Classification: dto.Info.Classification,
			Publisher:      dto.Info.Publisher,
			Price:          dto.Info.Price,
		},
		Tracklist: tracklist,
		Credits:   credits,
	}
}

func FullAlbumFromPb(pb *albumPb.FindFullAlbumResponse) FullAlbum {
	album := AlbumFromPb(pb.Album)
	info := InfoFromPb(pb.Info)

	pbCredits := pb.GetCredits()
	credits := make([]CreditInfo, len(pbCredits))
	for i := 0; i < len(pbCredits); i++ {
		credits[i] = CreditFromPb(pbCredits[i])
	}

	pbTracklist := pb.GetTracklist()
	tracklist := make([]Track, len(pbTracklist))
	for i := 0; i < len(pbTracklist); i++ {
		tracklist[i] = TrackFromPb(pbTracklist[i])
	}

	return FullAlbum{
		Album:     album,
		Info:      info,
		Credits:   credits,
		Tracklist: tracklist,
	}
}

func InfoFromPb(pb *albumPb.AlbumInfo) Info {
	return Info{
		CatalogNumber:  pb.GetCatalogNumber(),
		FullImageSrc:   pb.GetFullImageSrc(),
		SmallImageSrc:  pb.GetSmallImageSrc(),
		Barcode:        pb.GetBarcode(),
		CurrencyCode:   pb.GetCurrencyCode(),
		MediaFormat:    pb.GetMediaFormat(),
		Classification: pb.GetClassification(),
		Publisher:      pb.GetPublisher(),
		Price:          pb.GetPrice(),
	}
}

func CreditFromPb(pb *albumPb.CreditInfo) CreditInfo {
	return CreditInfo{
		FirstName:  pb.FirstName,
		LastName:   pb.LastName,
		Profession: pb.Profession,
	}
}

func TrackFromPb(pb *albumPb.TrackInfo) Track {
	return Track{
		ID:       pb.GetId(),
		AlbumID:  pb.GetAlbumId(),
		Title:    pb.GetTitle(),
		Duration: pb.GetDuration(),
	}
}

func PersonFromDto(dto dto.Person) Person {
	return Person{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		BirthDate: dto.BirthDate,
	}
}

func (t *Track) PbFromkModel() *albumPb.Track {
	return &albumPb.Track{
		Title:    t.Title,
		Duration: t.Duration,
	}
}

func (t *Track) DtoFromkModel() dto.TrackRes {
	return dto.TrackRes{
		ID:       t.ID,
		AlbumID:  t.AlbumID,
		Title:    t.Title,
		Duration: t.Duration,
	}
}

func (c *Credit) PbFromkModel() *albumPb.Credit {
	return &albumPb.Credit{
		Profession: c.Profession,
		PersonId:   c.PersonID,
	}
}

func (c *CreditInfo) DtoFromkModel() dto.CreditInfoRes {
	return dto.CreditInfoRes{
		Profession: c.Profession,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
	}
}

func (i *Info) DtoFromkModel() dto.InfoRes {
	return dto.InfoRes{
		CatalogNumber:  i.CatalogNumber,
		FullImageSrc:   i.FullImageSrc,
		SmallImageSrc:  i.SmallImageSrc,
		Barcode:        i.Barcode,
		CurrencyCode:   i.CurrencyCode,
		MediaFormat:    i.MediaFormat,
		Classification: i.Classification,
		Publisher:      i.Publisher,
		Price:          i.Price,
	}
}

func AlbumFromPb(pb *albumPb.Album) AlbumView {
	return AlbumView{
		ID:         pb.GetAlbumId(),
		Title:      pb.GetTitle(),
		CreatedAt:  pb.GetCreatedAt(),
		ReleasedAt: pb.GetReleasedAt(),
	}
}

func AlbumPreviewFromPb(pb *albumPb.AlbumPreview) AlbumPreview {
	return AlbumPreview{
		ID:            pb.GetAlbumId(),
		Title:         pb.GetTitle(),
		CreatedAt:     pb.GetCreatedAt(),
		ReleasedAt:    pb.GetReleasedAt(),
		Publisher:     pb.GetPublisher(),
		SmallImageSrc: pb.GetSmallImageSrc(),
	}
}

func (a *AlbumView) DtoFromModel() dto.AlbumViewRes {
	return dto.AlbumViewRes{
		AlbumID:    a.ID,
		Title:      a.Title,
		ReleasedAt: a.ReleasedAt,
		CreatedAt:  a.CreatedAt,
	}
}

func (a *AlbumPreview) DtoFromModel() dto.AlbumPreviewRes {
	return dto.AlbumPreviewRes{
		AlbumID:       a.ID,
		Title:         a.Title,
		ReleasedAt:    a.ReleasedAt,
		CreatedAt:     a.CreatedAt,
		Publisher:     a.Publisher,
		SmallImageSrc: a.SmallImageSrc,
	}
}

func (f *FullAlbum) DtoFromModel() dto.FullAlbumRes {
	album := f.Album.DtoFromModel()
	info := f.Info.DtoFromkModel()

	credits := make([]dto.CreditInfoRes, len(f.Credits))
	for i := 0; i < len(f.Credits); i++ {
		credits[i] = f.Credits[i].DtoFromkModel()
	}

	tracklist := make([]dto.TrackRes, len(f.Tracklist))
	for i := 0; i < len(f.Tracklist); i++ {
		tracklist[i] = f.Tracklist[i].DtoFromkModel()
	}

	return dto.FullAlbumRes{
		Album:     album,
		Info:      info,
		Credits:   credits,
		Tracklist: tracklist,
	}
}
