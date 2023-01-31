package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
)

type InfoService interface {
	GetAll(ctx context.Context) ([]*model.Info, error)
	Create(ctx context.Context, album model.Info) (model.Info, error)
	GetOne(ctx context.Context, albumID string) (model.Info, error)
}

type infoPolicy struct {
	infoService InfoService
}

func NewInfoPolicy(service InfoService) *infoPolicy {
	return &infoPolicy{infoService: service}
}

func (p *infoPolicy) GetAll(ctx context.Context) ([]*model.Info, error) {
	products, err := p.infoService.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "infoService.All")
	}

	return products, nil
}

func (p *infoPolicy) Create(ctx context.Context, info model.Info) (model.Info, error) {
	return p.infoService.Create(ctx, info)
}

func (p *infoPolicy) GetOne(ctx context.Context, albumID string) (model.Info, error) {
	return p.infoService.GetOne(ctx, albumID)
}
