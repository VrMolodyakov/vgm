package dao

import (
	"time"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	mapper "github.com/worldline-go/struct2"
)

type AlbumStorage struct {
	ID         string    `struct:"album_id,omitempty"`
	Title      string    `struct:"title,omitempty"`
	ReleasedAt time.Time `struct:"released_at,omitempty"`
	CreatedAt  time.Time `struct:"created_at,omitempty"`
}

func NewAlbumStorage(album model.Album) AlbumStorage {
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

func ToStorageMap(album model.Album) map[string]interface{} {
	storage := NewAlbumStorage(album)
	albumStorageMap := (&mapper.Decoder{}).Map(storage)
	return albumStorageMap
}

// type AlbumStorage struct {
// 	ID         string    `struct:"album_id"`
// 	Title      string    `struct:title,omitempty`
// 	ReleasedAt time.Time `struct:released_at,omitempty`
// 	CreatedAt  time.Time `struct:created_at,omitempty`
// }
