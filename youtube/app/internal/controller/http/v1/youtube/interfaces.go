package youtube

import "context"

type YoutubeService interface {
	GetVideoIDByTitle(ctx context.Context, videoTitle string) (string, error)
	CreatePlaylist(title string) (string, error)
	AddVideosToPlaylist(playlistID string, videoIDs []string) error
}

type MusicService interface {
	FindRandomTitles(ctx context.Context, count uint64) ([]string, error)
}
