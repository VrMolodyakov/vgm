package service

import (
	"context"
	"fmt"

	"google.golang.org/api/youtube/v3"
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

func (s *youtubeService) CreatePlaylist(title string) (string, error) {
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

func (s *youtubeService) AddVideosToPlaylist(playlistID string, videoIDs []string) error {
	for _, videoID := range videoIDs {
		playlistItem := &youtube.PlaylistItem{
			Snippet: &youtube.PlaylistItemSnippet{
				PlaylistId: playlistID,
				ResourceId: &youtube.ResourceId{
					Kind:    "youtube#video",
					VideoId: videoID,
				},
			},
		}

		_, err := s.youtube.PlaylistItems.Insert([]string{"snippet"}, playlistItem).Do()
		if err != nil {
			return fmt.Errorf("cannot add to playlist due to: %v", err)
		}
	}

	return nil
}
