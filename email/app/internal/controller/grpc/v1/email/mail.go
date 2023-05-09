package email

import (
	"context"

	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	"github.com/VrMolodyakov/vgm/email/app/internal/controller/grpc/v1/metrics"
	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/model"
	grpcerrors "github.com/VrMolodyakov/vgm/email/app/pkg/grpc-errors"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("email-grpc")
)

func (s *server) CreateEmail(ctx context.Context, req *emailPb.CreateEmailRequest) (*emailPb.CreateEmailResponse, error) {
	ctx, span := tracer.Start(ctx, "server.CreateEmail")
	defer span.End()

	email := model.ModelFromPB(req)
	if err := email.Validate(); err != nil {
		metrics.ErrorRequests.Inc()
		s.logger.Errorf("Validation %w", err)
		return &emailPb.CreateEmailResponse{},
			grpcerrors.ErrorResponse(err, err.Error())
	}
	if err := s.emailUseCae.Publush(ctx, email); err != nil {
		metrics.ErrorRequests.Inc()
		s.logger.Error(err)
		return &emailPb.CreateEmailResponse{},
			grpcerrors.ErrorResponse(err, err.Error())
	}
	metrics.SuccessRequests.Inc()
	return &emailPb.CreateEmailResponse{}, nil
}
