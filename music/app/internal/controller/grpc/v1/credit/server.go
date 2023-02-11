package credit

import (
	"context"

	creditPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/credit/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
)

type CreditPolicy interface {
	Create(ctx context.Context, credits []model.Credit) error
	GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error)
	Delete(ctx context.Context, albumID string) error
}

type server struct {
	creditPolicy CreditPolicy
	creditPb.UnimplementedCreditServiceServer
}

func NewServer(credit CreditPolicy, s creditPb.UnimplementedCreditServiceServer) *server {
	return &server{
		creditPolicy:                     credit,
		UnimplementedCreditServiceServer: s,
	}
}
