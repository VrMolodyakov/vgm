package email

import (
	"context"

	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/model"
	grpcerrors "github.com/VrMolodyakov/vgm/email/app/pkg/grpc-errors"
)

//TODO:publish context
func (s *server) Create(req *emailPb.CreateEmailRequest) (*emailPb.CreateEmailResponse, error) {
	email := model.ModelFromPB(req)
	if err := email.Validate(); err != nil {
		s.logger.Errorf("Validation %w", err)
		return &emailPb.CreateEmailResponse{},
			grpcerrors.ErrorResponse(err, err.Error())
	}
	if err := s.emailUseCae.Publush(context.Background(), email); err != nil {
		s.logger.Error(err)
		return &emailPb.CreateEmailResponse{},
			grpcerrors.ErrorResponse(err, err.Error())
	}
	return &emailPb.CreateEmailResponse{}, nil
}
