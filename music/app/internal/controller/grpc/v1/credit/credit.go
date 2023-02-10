package credit

import (
	_ "context"
	// _creditPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/credit/v1"
	// _creditModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
)

// func (s *server) CreateCredit(ctx context.Context, request *creditPb.CreateCreditRequest) (*creditPb.CreateCreditResponse, error) {
// 	credit := creditModel.NewCreditFromPB(request)
// 	profID, err := s.profPolicy.GetOne(ctx, request.GetProfession())
// 	if err != nil {
// 		profID, err = s.profPolicy.Create(ctx, request.GetProfession())
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	credit.Profession = profID.ID
// 	credit, err = s.creditPolicy.Create(ctx, credit)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &creditPb.CreateCreditResponse{}, nil
// }

// func (s *server) FindCredits(ctx context.Context, request *creditPb.FindCreditsRequest) (*creditPb.FindCreditsResponse, error) {
// 	creditsModel, err := s.creditPolicy.GetAll(ctx, request.GetAlbumId())
// 	if err != nil {
// 		return nil, err
// 	}
// 	credits := make([]*creditPb.Credit, len(creditsModel))
// 	for i := 0; i < len(creditsModel); i++ {
// 		credits[i] = creditsModel[i].ToProto()
// 	}
// 	return &creditPb.FindCreditsResponse{Credits: credits}, nil
// }
