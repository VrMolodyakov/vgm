package album

import (
	"context"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
)

func (s *server) CreateAlbum(context.Context, *albumPb.CreateAlbumRequest) (*albumPb.CreateAlbumResponse, error) {
	return nil, nil
}
func (s *server) FindAlbum(context.Context, *albumPb.FindAlbumRequest) (*albumPb.FindAlbumResponse, error) {
	return nil, nil
}
func (s *server) FindAllAlbums(context.Context, *albumPb.FindAllAlbumsRequest) (*albumPb.FindAllAlbumsResponse, error) {
	return nil, nil
}
func (s *server) FindFullAlbum(context.Context, *albumPb.FindFullAlbumRequest) (*albumPb.FindFullAlbumResponse, error) {
	return nil, nil
}
