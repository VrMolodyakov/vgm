package youtube

import (
	"context"
	"time"
)

type YoutubeService interface {
	GetVideoIDByTitle(ctx context.Context, videoTitle string) (string, error)
	CreatePlaylist(ctx context.Context, title string) (string, error)
	AddVideosToPlaylist(ctx context.Context, playlistID string, videoIDs []string) error
}

type MusicService interface {
	FindRandomTitles(ctx context.Context, count uint64) ([]string, error)
}

type Cache interface {
	Set(key string, value string, expireAt time.Duration) string
	Get(key string) (string, bool)
}
