package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
)

type CreditService interface {
	Create(ctx context.Context, credit model.Credit) (model.Credit, error)
	GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error)
}

type CreditPolicy struct {
	creditService CreditService
}

func NewCreditPolicy(service CreditService) *CreditPolicy {
	return &CreditPolicy{creditService: service}
}

func (p *CreditPolicy) Create(ctx context.Context, credit model.Credit) (model.Credit, error) {
	return p.creditService.Create(ctx, credit)
}

func (p *CreditPolicy) GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error) {
	return p.creditService.GetAll(ctx, albumID)
}
