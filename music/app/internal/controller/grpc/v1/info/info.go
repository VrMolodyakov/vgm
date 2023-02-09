package info

import (
	"context"

	infoPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/info/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
)

func (s *server) FindAlbumInfo(ctx context.Context, request *infoPb.FindAlbumInfoRequest) (*infoPb.FindAlbumInfoResponse, error) {
	id := request.GetAlbumId()
	info, err := s.infoPolicy.GetOne(ctx, id)

	if err != nil {
		return nil, err
	}

	return &infoPb.FindAlbumInfoResponse{
		Info: info.ToProto(),
	}, nil

}

func (s *server) UpdateAlbumInfo(ctx context.Context, request *infoPb.UpdateAlbumInfoRequest) (*infoPb.UpdateAlbumInfoResponse, error) {
	info := model.UpdateModelFromPB(request)
	err := s.infoPolicy.Update(ctx, info)
	if err != nil {
		return nil, err
	}

	return &infoPb.UpdateAlbumInfoResponse{}, nil
}

func (s *server) DeleteAlbumInfo(ctx context.Context, request *infoPb.DeleteAlbumInfoRequest) (*infoPb.DeleteAlbumInfoResponse, error) {
	id := request.GetAlbumId()
	if id == "" {
		id = request.GetAlbumInfoId()
	}
	err := s.infoPolicy.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return &infoPb.DeleteAlbumInfoResponse{}, nil
}
