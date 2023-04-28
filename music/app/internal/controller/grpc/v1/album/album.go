package album

import (
	"context"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	albumModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	personModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	tracer = otel.Tracer("music-server")
)

func (s *server) FindAllAlbums(ctx context.Context, request *albumPb.FindAllAlbumsRequest) (*albumPb.FindAllAlbumsResponse, error) {
	ctx, span := tracer.Start(ctx, "music-server.FindAll")
	defer span.End()
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
		album := a.ToProto()
		pbAlbums[i] = &album
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
	ctx, span := tracer.Start(ctx, "music-server.FindFullAlbum")
	defer span.End()
	fullAlbum, err := s.albumPolicy.GetOne(ctx, request.GetAlbumId())
	if err != nil {
		return &albumPb.FindFullAlbumResponse{}, err
	}
	info := fullAlbum.Info.ToProto()
	album := fullAlbum.Album.ToProto()

	tracklist := []*albumPb.TrackInfo{}

	for i := 0; i < len(fullAlbum.Tracklist); i++ {
		track := fullAlbum.Tracklist[i].ToProto()
		tracklist = append(tracklist, &track)
	}

	credits := []*albumPb.CreditInfo{}

	for i := 0; i < len(fullAlbum.Credits); i++ {
		credit := fullAlbum.Credits[i].ToProto()
		credits = append(credits, &credit)
	}

	return &albumPb.FindFullAlbumResponse{
		Album:     &album,
		Info:      &info,
		Credits:   credits,
		Tracklist: tracklist,
	}, nil
}

func (s *server) FindAlbum(context.Context, *albumPb.FindAlbumRequest) (*albumPb.FindAlbumResponse, error) {
	return nil, nil
}

func (s *server) CreatePerson(ctx context.Context, request *albumPb.CreatePersonRequest) (*albumPb.CreatePersonResponse, error) {
	personModel := personModel.NewPersonFromPB(request)
	person, err := s.personPolicy.Create(ctx, personModel)

	if err != nil {
		return nil, err
	}

	return &albumPb.CreatePersonResponse{
		Person: person.ToProto(),
	}, nil
}
