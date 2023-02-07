package album

import (
	"context"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	trackPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/track/v1"
	albumModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	infoModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	trackModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
)

// func (s *server) CreateAlbum(ctx context.Context, request *albumPb.CreateAlbumRequest) (*albumPb.CreateAlbumResponse, error) {
// 	a := albumModel.NewAlbumFromPB(request)

// 	album, err := s.albumPolicy.Create(ctx, a)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &albumPb.CreateAlbumResponse{
// 		Album: album.ToProto(),
// 	}, nil
// }

func (s *server) FindAllAlbums(ctx context.Context, request *albumPb.FindAllAlbumsRequest) (*albumPb.FindAllAlbumsResponse, error) {
	sort := albumModel.AlbumSort(request)
	filter := albumModel.AlbumFilter(request)

	all, err := s.albumPolicy.GetAll(ctx, filter, sort)
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

func (s *server) CreateFullAlbum(ctx context.Context, request *albumPb.CreateFullAlbumRequest) (*albumPb.CreateFullAlbumResponse, error) {
	albumModel := albumModel.NewAlbumFromPB(request)
	infoModel := infoModel.NewInfoFromPB(request)

	album, err := s.albumPolicy.Create(ctx, albumModel)
	if err != nil {
		return nil, err
	}
	infoModel.AlbumID = album.ID
	info, err := s.infoPolicy.Create(ctx, infoModel)

	if err != nil {
		return nil, err
	}
	tracklistPb := request.GetTracklist()
	tracklist := make([]trackModel.Track, len(tracklistPb))
	for i := 0; i < len(tracklistPb); i++ {
		tracklist[i] = trackModel.NewTrackFromPB(tracklistPb[i])
		tracklist[i].AlbumID = album.ID
	}
	err = s.trackPolicy.Create(ctx, tracklist)
	if err != nil {
		return nil, err
	}

	protoAlbum := album.ToProto()
	protoInfo := info.ToProto()
	protoTracklist := make([]*trackPb.Track, len(tracklist))
	for i := 0; i < len(tracklist); i++ {
		protoTracklist[i] = tracklist[i].ToProto()
	}

	return &albumPb.CreateFullAlbumResponse{
		Album:     protoAlbum,
		Info:      protoInfo,
		Tracklist: protoTracklist,
	}, nil
}

func (s *server) FindFullAlbum(ctx context.Context, request *albumPb.FindFullAlbumRequest) (*albumPb.FindFullAlbumResponse, error) {
	return nil, nil
}

func (s *server) FindAlbum(ctx context.Context, request *albumPb.FindAlbumRequest) (*albumPb.FindAlbumResponse, error) {
	return nil, nil
}
