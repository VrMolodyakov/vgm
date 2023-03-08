package client

import (
	"log"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/album/model"
	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type musicClient struct {
	target string
	albumPb.AlbumServiceClient
}

func NewMusicClient(target string) *musicClient {
	if target == "" {
		log.Fatalln("Error in Access to GRPC URL in music client")
	}
	return &musicClient{
		target: target,
	}
}

func (m *musicClient) Start() {

	conn, err := grpc.Dial(m.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	m.AlbumServiceClient = albumPb.NewAlbumServiceClient(conn)
}

func (m *musicClient) Create(album model.Album) error {
	return nil
}
