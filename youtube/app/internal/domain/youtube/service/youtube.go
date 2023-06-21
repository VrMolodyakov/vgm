package service

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/youtube/v3"
)

var (
	tracer = otel.Tracer("youtube-api")
)

type YoutubeAPI interface {
	SearchListCall() *youtube.SearchListCall
}

type youtubeService struct {
	youtube *youtube.Service
}

func NewYoutubeService(youtube *youtube.Service) *youtubeService {
	return &youtubeService{
		youtube: youtube,
	}
}

func (s *youtubeService) GetVideoIDByTitle(ctx context.Context, videoTitle string) (string, error) {
	_, span := tracer.Start(ctx, "api.GetVideoIDByTitle")
	defer span.End()

	searchResponse, err := s.youtube.Search.List([]string{"id"}).Q(videoTitle).MaxResults(1).Do()
	if err != nil {
		return "", fmt.Errorf("error executing search query: %v", err)
	}

	for _, searchResult := range searchResponse.Items {
		if searchResult.Id.Kind == "youtube#video" {
			return searchResult.Id.VideoId, nil
		}
	}

	return "", nil
}

func (s *youtubeService) CreatePlaylist(ctx context.Context, title string) (string, error) {
	_, span := tracer.Start(ctx, "api.CreatePlaylist")
	defer span.End()

	playlist := &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title: title,
		},
		Status: &youtube.PlaylistStatus{
			PrivacyStatus: "public",
		},
	}

	playlistResponse, err := s.youtube.Playlists.Insert([]string{"snippet,status"}, playlist).Do()
	if err != nil {
		return "", fmt.Errorf("create playlist error: %v", err)
	}

	playlistID := playlistResponse.Id
	playlistURL := fmt.Sprintf("https://www.youtube.com/playlist?list=%s", playlistID)

	return playlistURL, nil
}

func (s *youtubeService) AddVideosToPlaylist(ctx context.Context, playlistID string, videoIDs []string) error {
	ctx, span := tracer.Start(ctx, "api.AddVideosToPlaylist")
	defer span.End()

	g, _ := errgroup.WithContext(ctx)
	for _, id := range videoIDs {
		id := id
		g.Go(func() error {
			playlistItem := &youtube.PlaylistItem{
				Snippet: &youtube.PlaylistItemSnippet{
					PlaylistId: playlistID,
					ResourceId: &youtube.ResourceId{
						Kind:    "youtube#video",
						VideoId: id,
					},
				},
			}

			_, err := s.youtube.PlaylistItems.Insert([]string{"snippet"}, playlistItem).Do()
			if err != nil {
				return fmt.Errorf("cannot add to playlist due to: %v", err)
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error when executing the gorutin group: %v", err)
	}

	return nil
}
