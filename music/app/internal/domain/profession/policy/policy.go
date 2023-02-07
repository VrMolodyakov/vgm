package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/profession/model"
)

type ProfessionService interface {
	Create(ctx context.Context, profession string) (model.Profession, error)
	GetOne(ctx context.Context, profession string) (model.Profession, error)
}

type professionPolicy struct {
	professionService ProfessionService
}

func NewProfessionPolicy(service ProfessionService) *professionPolicy {
	return &professionPolicy{professionService: service}
}

func (p *professionPolicy) GetOne(ctx context.Context, prof string) (model.Profession, error) {
	return p.professionService.GetOne(ctx, prof)

}

func (p *professionPolicy) Create(ctx context.Context, profession string) (model.Profession, error) {
	return p.professionService.Create(ctx, profession)
}
