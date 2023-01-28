package dao

import (
	"time"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
)

type AlbumStorage struct {
	ID         string
	Title      string
	ReleasedAt time.Time
	CreatedAt  time.Time
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
		CreatedAt:  album.CreatedAt.UnixMicro(),
		ReleasedAt: album.ReleasedAt.UnixMicro(),
	}
}
