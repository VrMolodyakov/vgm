package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("credit-policy")
)

type CreditService interface {
	Create(ctx context.Context, credits []model.Credit) error
	GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error)
	Delete(ctx context.Context, albumID string) error
	Update(ctx context.Context, albumId string, role string) error
}

type CreditPolicy struct {
	creditService CreditService
}

func NewCreditPolicy(service CreditService) *CreditPolicy {
	return &CreditPolicy{creditService: service}
}

func (p *CreditPolicy) Create(ctx context.Context, credits []model.Credit) error {
	_, span := tracer.Start(ctx, "policy.Create")
	defer span.End()
	return p.creditService.Create(ctx, credits)
}

func (p *CreditPolicy) GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error) {
	_, span := tracer.Start(ctx, "policy.GetAll")
	defer span.End()
	return p.creditService.GetAll(ctx, albumID)
}

func (p *CreditPolicy) Delete(ctx context.Context, albumID string) error {
	_, span := tracer.Start(ctx, "policy.Delete")
	defer span.End()
	return p.creditService.Delete(ctx, albumID)
}

func (p *CreditPolicy) Update(ctx context.Context, albumId string, role string) error {
	_, span := tracer.Start(ctx, "policy.Update")
	defer span.End()
	return p.creditService.Update(ctx, albumId, role)
}
