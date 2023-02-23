package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
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
	return p.creditService.Create(ctx, credits)
}

func (p *CreditPolicy) GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error) {
	return p.creditService.GetAll(ctx, albumID)
}

func (p *CreditPolicy) Delete(ctx context.Context, albumID string) error {
	return p.creditService.Delete(ctx, albumID)
}

func (p *CreditPolicy) Update(ctx context.Context, albumId string, role string) error {
	return p.creditService.Update(ctx, albumId, role)
}
