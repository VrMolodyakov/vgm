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

func (c *Credit) PbFromkModel() *albumPb.Credit {
	return &albumPb.Credit{
		Profession: c.Profession,
		PersonId:   c.PersonID,
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
