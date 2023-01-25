package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumService interface {
	All(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.Album, error)
	Create(ctx context.Context, album model.Album) (model.Album, error)
}

type albumPolicy struct {
	albumService AlbumService
}

func NewAlbumPolicy(service AlbumService) *albumPolicy {
	return &albumPolicy{albumService: service}
}

func (p *albumPolicy) All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.Album, error) {
	products, err := p.albumService.All(ctx, filtering, sorting)
	if err != nil {
		return nil, errors.Wrap(err, "albumService.All")
	}

	return products, nil
}

func (p *albumPolicy) CreateProduct(ctx context.Context, product model.Album) (model.Album, error) {
	return p.albumService.Create(ctx, product)
}
