package youtube

import "context"

type YoutubeService interface {
	GetVideoIDByTitle(ctx context.Context, videoTitle string) (string, error)
}

type MusicService interface {
	FindRandomTitles(ctx context.Context, count uint64) ([]string, error)
}
