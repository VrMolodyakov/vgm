package dao

import (
	"time"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	mapper "github.com/worldline-go/struct2"
)

const (
	fields = 3
)

type AlbumStorage struct {
	ID         string    `struct:"album_id"`
	Title      string    `struct:"title"`
	ReleasedAt time.Time `struct:"released_at"`
	CreatedAt  time.Time `struct:"created_at"`
}

func fromModel(album model.Album) AlbumStorage {
	createdAt := time.UnixMilli(album.CreatedAt)
	releasedAt := time.UnixMilli(album.ReleasedAt)
	return AlbumStorage{
		ID:         album.ID,
		Title:      album.Title,
		CreatedAt:  createdAt,
		ReleasedAt: releasedAt,
	}
}

func (album AlbumStorage) ToModel() model.Album {
	return model.Album{
		ID:         album.ID,
		Title:      album.Title,
		CreatedAt:  album.CreatedAt.UnixMilli(),
		ReleasedAt: album.ReleasedAt.UnixMilli(),
	}
}

func toStorageMap(album model.Album) map[string]interface{} {
	storage := fromModel(album)
	albumStorageMap := (&mapper.Decoder{}).Map(storage)
	return albumStorageMap
}

func ToUpdateStorageMap(album model.Album) map[string]interface{} {
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
