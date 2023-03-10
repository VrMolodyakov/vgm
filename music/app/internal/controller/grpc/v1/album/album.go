package album

import (
	"context"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	albumModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) FindAllAlbums(ctx context.Context, request *albumPb.FindAllAlbumsRequest) (*albumPb.FindAllAlbumsResponse, error) {
	sort := albumModel.AlbumSort(request)
	filter := albumModel.AlbumFilter(request)

	all, err := s.albumPolicy.GetAll(ctx, filter, sort)
	if err != nil {
		if _, ok := errors.IsInternal(err); ok {
			status.Error(codes.Internal, "internal server error")
		}
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

func (s *server) DeleteAlbum(ctx context.Context, request *albumPb.DeleteAlbumRequest) (*albumPb.DeleteAlbumResponse, error) {
	err := s.albumPolicy.Delete(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return &albumPb.DeleteAlbumResponse{}, nil
}

func (s *server) UpdateAlbum(ctx context.Context, request *albumPb.UpdateAlbumRequest) (*albumPb.UpdateAlbumResponse, error) {
	album := albumModel.UpdateModelFromPB(request)
	err := s.albumPolicy.Update(ctx, album)
	if err != nil {
		return nil, err
	}
	return &albumPb.UpdateAlbumResponse{}, nil
}

func (s *server) CreateAlbum(ctx context.Context, request *albumPb.CreateAlbumRequest) (*albumPb.CreateAlbumResponse, error) {
	album := albumModel.NewAlbumFromPB(request)
	err := s.albumPolicy.Create(ctx, *album)
	if err != nil {
		if _, ok := errors.IsInternal(err); ok {
			status.Error(codes.Internal, "internal server error")
		}
		return nil, err
	}
	return &albumPb.CreateAlbumResponse{}, nil

}

func (s *server) FindFullAlbum(ctx context.Context, request *albumPb.FindFullAlbumRequest) (*albumPb.FindFullAlbumResponse, error) {
	// album, err := s.albumPolicy.GetOne(context.Background(), request.GetAlbumId())
	// if err != nil {
	// 	return &albumPb.FindFullAlbumResponse{}, err
	// }
	// return &albumPb.FindFullAlbumResponse{
	// 	Album: album.Album.ToProto(),
	// 	Credits: album.,
	// }
	return nil, nil
}

func (s *server) FindAlbum(context.Context, *albumPb.FindAlbumRequest) (*albumPb.FindAlbumResponse, error) {
	return nil, nil
}
