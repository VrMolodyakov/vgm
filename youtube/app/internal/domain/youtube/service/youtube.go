package service

import (
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

func (s *youtubeService) GetVideoIDByTitle(videoTitle string) (string, error) {
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
