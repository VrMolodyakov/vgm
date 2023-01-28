package dao

import (
	"time"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	mapper "github.com/worldline-go/struct2"
)

type AlbumStorage struct {
	ID         string    `struct:"album_id"`
	Title      string    `struct:"title"`
	ReleasedAt time.Time `struct:"released_at"`
	CreatedAt  time.Time `struct:"created_at"`
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

func ToStorageMap(album model.Album) (map[string]interface{}, error) {
	storage := NewAlbumStorage(album)
	logger := logging.GetLogger()
	logger.Info("-----STORAGE STRUCT-----")
	logger.Sugar().Info(storage)

	albumStorageMap := (&mapper.Decoder{}).Map(storage)
	// var updateAlbumStorageMap map[string]interface{}
	// err := mapstructure.Decode(&storage, &updateAlbumStorageMap)
	// if err != nil {
	// 	return updateAlbumStorageMap, errors.Wrap(err, "mapstructure.Decode(product)")
	// }
	logger.Info("-----STRUCT AS MAP-----")
	logger.Sugar().Info(albumStorageMap)
	return albumStorageMap, nil
}
