package album

import (
	"context"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
)

func (s *server) CreateAlbum(ctx context.Context, request *albumPb.CreateAlbumRequest) (*albumPb.CreateAlbumResponse, error) {
	a := model.NewAlbumFromPB(request)

	album, err := s.albumPolicy.CreateAlbum(ctx, a)
	if err != nil {
		return nil, err
	}

	return &albumPb.CreateAlbumResponse{
		Album: album.ToProto(),
	}, nil
}

func (s *server) FindAllAlbums(ctx context.Context, request *albumPb.FindAllAlbumsRequest) (*albumPb.FindAllAlbumsResponse, error) {
	sort := model.AlbumSort(request)
	filter := model.AlbumFilter(request)

	all, err := s.albumPolicy.All(ctx, filter, sort)
	if err != nil {
		return nil, err
	}

	pbAlbums := make([]*albumPb.Album, len(all))
	for i, a := range all {
		pbAlbums[i] = a.ToProto()
	}

	return &albumPb.FindAllAlbumsResponse{
		Albums: pbAlbums,
	}, nil
}

func (s *server) FindFullAlbum(ctx context.Context, request *albumPb.FindFullAlbumRequest) (*albumPb.FindFullAlbumResponse, error) {
	return nil, nil
}

func (s *server) FindAlbum(ctx context.Context, request *albumPb.FindAlbumRequest) (*albumPb.FindAlbumResponse, error) {
	return nil, nil
}
