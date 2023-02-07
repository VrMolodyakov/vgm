package credit

import (
	"context"

	creditPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/credit/v1"
	creditModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
)

func (s *server) CreateCredit(ctx context.Context, request *creditPb.CreateCreditRequest) (*creditPb.CreateCreditResponse, error) {
	credit := creditModel.NewCreditFromPB(request)
	profID, err := s.profPolicy.GetOne(ctx, request.GetProfession())
	if err != nil {
		profID, err = s.profPolicy.Create(ctx, request.GetProfession())
		if err != nil {
			return nil, err
		}
	}
	credit.ProfessionID = profID.ID
	credit, err = s.creditPolicy.Create(ctx, credit)
	if err != nil {
		return nil, err
	}
	return &creditPb.CreateCreditResponse{}, nil
}
