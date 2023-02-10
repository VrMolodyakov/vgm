package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	creditModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	infoModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	trackModel "github.com/VrMolodyakov/vgm/music/app/internal/domain/tracklist/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumService interface {
	GetAll(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.Album, error)
	Create(ctx context.Context, album model.Album) (model.Album, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.Album) error
}

type InfoService interface {
	Create(ctx context.Context, Info infoModel.Info) (infoModel.Info, error)
}

type TrackService interface {
	Create(ctx context.Context, tracklist []trackModel.Track) error
}

type CreditService interface {
	Create(ctx context.Context, credit creditModel.Credit) (creditModel.Credit, error)
}

type albumPolicy struct {
	albumService  AlbumService
	infoService   InfoService
	trackService  TrackService
	creditService CreditService
}

func NewAlbumPolicy(service AlbumService) *albumPolicy {
	return &albumPolicy{albumService: service}
}

func (p *albumPolicy) GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.Album, error) {
	products, err := p.albumService.GetAll(ctx, filtering, sorting)
	if err != nil {
		return nil, errors.Wrap(err, "albumService.All")
	}

	return products, nil
}

func (p *albumPolicy) Create(ctx context.Context, album model.Album) (model.Album, error) {
	return p.albumService.Create(ctx, album)
}

func (p *albumPolicy) Delete(ctx context.Context, id string) error {
	return p.albumService.Delete(ctx, id)
}

func (p *albumPolicy) Update(ctx context.Context, album model.Album) error {
	return p.albumService.Update(ctx, album)
}

func (p *albumPolicy) Create2(ctx context.Context, album model.FullAlbum) (*model.FullAlbum, error) {
	albumModel, err := p.albumService.Create(ctx, album.Album)
	if err != nil {
		return nil, err
	}

	infoModel, err := p.infoService.Create(ctx, album.Info)
	if err != nil {
		return nil, err
	}

	err = p.trackService.Create(ctx, album.Tracklist)
	if err != nil {
		return nil, err
	}
	p.creditService.Create(ctx, album.Credits)

	return &model.FullAlbum{}, nil
}
