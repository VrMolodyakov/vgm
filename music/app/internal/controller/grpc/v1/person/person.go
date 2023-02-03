package person

import (
	"context"

	personPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/person/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
)

func (s *server) CreatePerson(ctx context.Context, request *personPb.CreatePersonRequest) (*personPb.CreatePersonResponse, error) {
	personModel := model.NewAlbumFromPB(request)
	person, err := s.personPolicy.Create(ctx, personModel)

	if err != nil {
		return nil, err
	}

	return &personPb.CreatePersonResponse{
		Person: person.ToProto(),
	}, nil
}

func (s *server) FindAllPersons(ctx context.Context, request *personPb.FindAllPersonsRequest) (*personPb.FindAllPersonsResponse, error) {
	filter := model.PersonFilter(request)

	all, err := s.personPolicy.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	pbPersons := make([]*personPb.Person, len(all))
	for i, p := range all {
		pbPersons[i] = p.ToProto()
	}

	return &personPb.FindAllPersonsResponse{
		Person: pbPersons,
	}, nil
}
