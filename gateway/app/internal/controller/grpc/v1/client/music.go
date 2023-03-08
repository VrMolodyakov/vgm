package client

import (
	"log"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"google.golang.org/grpc"
)

type musicClient struct {
	target string
	albumPb.AlbumServiceClient
}

func NewMusicClient(target string) *musicClient {
	return &musicClient{
		target: target,
	}
}

func (m *musicClient) starMusicClient() {

	conn, err := grpc.Dial(m.target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	m.AlbumServiceClient = albumPb.NewAlbumServiceClient(conn)
}
