package main

import (
	"context"
	"fmt"
	"log"

	"github.com/VrMolodyakov/vgm/youtube/internal/domain/youtube/service"
	"github.com/VrMolodyakov/vgm/youtube/pkg/youtube"
)

func main() {
	ctx := context.Background()
	client, err := youtube.NewYoutubeClient(ctx, "AIzaSyA0lomsxxv-ffFkukOcA5QINXtIOF0kEUc")
	if err != nil {
		log.Fatal(err)
	}
	s := service.NewYoutubeService(client)
	id, err := s.GetVideoIDByTitle("street fighter 3 album")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
