package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
)

type InfoService interface {
	Create(ctx context.Context, Info model.Info) (model.Info, error)
	GetOne(ctx context.Context, albumID string) (model.Info, error)
	Update(ctx context.Context, info model.Info) error
	Delete(ctx context.Context, id string) error
}

type infoPolicy struct {
	infoService InfoService
}

func NewInfoPolicy(service InfoService) *infoPolicy {
	return &infoPolicy{infoService: service}
}

func (p *infoPolicy) Create(ctx context.Context, info model.Info) (model.Info, error) {
	return p.infoService.Create(ctx, info)
}

func (p *infoPolicy) GetOne(ctx context.Context, albumID string) (model.Info, error) {
	return p.infoService.GetOne(ctx, albumID)
}

func (p *infoPolicy) Update(ctx context.Context, info model.Info) error {
	return p.infoService.Update(ctx, info)
}
func (p *infoPolicy) Delete(ctx context.Context, id string) error {
	return p.infoService.Delete(ctx, id)
}
