package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/profession/model"
)

type ProfessionService interface {
	Create(ctx context.Context, profession model.Profession) (model.Profession, error)
	GetOne(ctx context.Context, profID string) (model.Profession, error)
}

type professionPolicy struct {
	professionService ProfessionService
}

func NewProfessionPolicy(service ProfessionService) *professionPolicy {
	return &professionPolicy{professionService: service}
}

func (p *professionPolicy) GetOne(ctx context.Context, profID string) (model.Profession, error) {
	return p.professionService.GetOne(ctx, profID)

}

func (p *professionPolicy) Create(ctx context.Context, Profession model.Profession) (model.Profession, error) {
	return p.professionService.Create(ctx, Profession)
}
