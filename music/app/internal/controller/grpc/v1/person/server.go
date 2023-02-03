package person

import (
	"context"

	personPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/person/v1"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
)

type PersonPolicy interface {
	GetAll(ctx context.Context, filter filter.Filterable) ([]model.Person, error)
	Create(ctx context.Context, P model.Person) (model.Person, error)
}

type server struct {
	personPolicy PersonPolicy
	personPb.UnimplementedPersonServiceServer
}

func NewServer(policy PersonPolicy, s personPb.UnimplementedPersonServiceServer) *server {
	return &server{
		personPolicy:                     policy,
		UnimplementedPersonServiceServer: s,
	}
}
