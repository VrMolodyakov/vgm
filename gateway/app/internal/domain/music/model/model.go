package model

import (
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/handler/album/dto"
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

func AlbumFromDto(dto dto.Album) Album {
	tracklist := make([]Track, len(dto.Tracklist))
	credits := make([]Credit, len(dto.Credits))

	for i := 0; i < len(tracklist); i++ {
		tracklist[i] = Track{
			dto.Tracklist[i].ID,
			dto.Tracklist[i].AlbumID,
			dto.Tracklist[i].Title,
			dto.Tracklist[i].Duration,
		}
	}

	for i := 0; i < len(credits); i++ {
		credits[i] = Credit{
			dto.Credits[i].PersonID,
			dto.Credits[i].AlbumID,
			dto.Credits[i].Profession,
		}
	}

	return Album{
		Album: AlbumView{
			dto.Album.ID,
			dto.Album.Title,
			dto.Album.ReleasedAt,
			dto.Album.CreatedAt,
		},
		Info: Info{
			dto.Info.ID,
			dto.Info.AlbumID,
			dto.Info.CatalogNumber,
			dto.Info.FullImageSrc,
			dto.Info.SmallImageSrc,
			dto.Info.Barcode,
			dto.Info.CurrencyCode,
			dto.Info.MediaFormat,
			dto.Info.Classification,
			dto.Info.Publisher,
			dto.Info.Price,
		},
		Tracklist: tracklist,
		Credits:   credits,
	}
}

func FullAlbumFromPb(pb *albumPb.FindFullAlbumResponse) FullAlbum {
	album := AlbumFromPb(pb.Album)
	info := InfoFromPb(pb.Info)

	credits := make([]CreditInfo, len(pb.Credits))
	for i := 0; i < len(pb.Credits); i++ {
		credits[i] = CreditFromPb(pb.Credits[i])
	}

	tracklist := make([]Track, len(pb.Tracklist))
	for i := 0; i < len(pb.Credits); i++ {
		tracklist[i] = TrackFromPb(pb.Tracklist[i])
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
		CatalogNumber:  pb.CatalogNumber,
		FullImageSrc:   *pb.FullImageSrc,
		SmallImageSrc:  *pb.SmallImageSrc,
		Barcode:        *pb.Barcode,
		CurrencyCode:   pb.CurrencyCode,
		MediaFormat:    pb.MediaFormat,
		Classification: pb.Classification,
		Publisher:      pb.Publisher,
		Price:          pb.Price,
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
		ID:       pb.Id,
		AlbumID:  pb.AlbumId,
		Title:    pb.Title,
		Duration: pb.Duration,
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

func (t *Track) DtoFromkModel() dto.Track {
	return dto.Track{
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

func (c *CreditInfo) DtoFromkModel() dto.CreditInfo {
	return dto.CreditInfo{
		Profession: c.Profession,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
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

func (a *AlbumView) DtoFromModel() dto.AlbumView {
	return dto.AlbumView{
		ID:         a.ID,
		Title:      a.Title,
		ReleasedAt: a.ReleasedAt,
		CreatedAt:  a.CreatedAt,
	}
}

func (f *FullAlbum) DtoFromModel() dto.FullAlbumResponse {
	album := dto.AlbumView{
		ID:         f.Album.ID,
		Title:      f.Album.Title,
		ReleasedAt: f.Album.ReleasedAt,
		CreatedAt:  f.Album.CreatedAt,
	}
	info := dto.Info{
		ID:             f.Info.ID,
		AlbumID:        f.Info.AlbumID,
		CatalogNumber:  f.Info.CatalogNumber,
		FullImageSrc:   f.Info.FullImageSrc,
		SmallImageSrc:  f.Info.SmallImageSrc,
		Barcode:        f.Info.Barcode,
		CurrencyCode:   f.Info.CurrencyCode,
		MediaFormat:    f.Info.MediaFormat,
		Classification: f.Info.Classification,
		Publisher:      f.Info.Publisher,
		Price:          f.Info.Price,
	}

	credits := make([]dto.CreditInfo, len(f.Credits))
	for i := 0; i < len(f.Credits); i++ {
		credits[i] = f.Credits[i].DtoFromkModel()
	}

	tracklist := make([]dto.Track, len(f.Tracklist))
	for i := 0; i < len(f.Tracklist); i++ {
		tracklist[i] = f.Tracklist[i].DtoFromkModel()
	}

	return dto.FullAlbumResponse{
		Album:     album,
		Info:      info,
		Credits:   credits,
		Tracklist: tracklist,
	}
}
