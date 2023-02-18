package reposotory

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	mapper "github.com/worldline-go/struct2"
)

type TrackStorage struct {
	ID       int64  `struct:"track_id "`
	AlbumID  string `struct:"album_id"`
	Title    string `struct:"title"`
	Duration string `struct:"duration"`
}

func toStorageMap(tracklist []model.Track) map[string]interface{} {
	storageList := make([]TrackStorage, len(tracklist))
	for i := 0; i < len(tracklist); i++ {
		storageList[i] = fromModel(tracklist[i])
	}
	storageTracklist := (&mapper.Decoder{}).Map(storageList)
	return storageTracklist
}

func fromModel(track model.Track) TrackStorage {
	return TrackStorage{
		ID:      track.ID,
		AlbumID: track.AlbumID,
		Title:   track.Title,
	}
}

func (t *TrackStorage) toModel() model.Track {
	return model.Track{
		ID:      t.ID,
		AlbumID: t.AlbumID,
		Title:   t.Title,
	}
}
