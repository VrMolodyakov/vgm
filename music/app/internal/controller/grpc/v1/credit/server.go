package credit

import (
	"context"

	creditPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/credit/v1"
	creditModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	profModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/profession/model"
)

type CreditPolicy interface {
	Create(ctx context.Context, credit creditModel.Credit) (creditModel.Credit, error)
	GetAll(ctx context.Context, albumID string) ([]creditModel.CreditInfo, error)
}

type ProfessionPolicy interface {
	Create(ctx context.Context, profession string) (profModel.Profession, error)
	GetOne(ctx context.Context, profession string) (profModel.Profession, error)
}

type server struct {
	creditPolicy CreditPolicy
	profPolicy   ProfessionPolicy
	creditPb.UnimplementedCreditServiceServer
}

func NewServer(policy CreditPolicy, s creditPb.UnimplementedCreditServiceServer) *server {
	return &server{
		creditPolicy:                     policy,
		UnimplementedCreditServiceServer: s,
	}
}
