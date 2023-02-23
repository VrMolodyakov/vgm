package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type albumPolicy struct {
	albumService  AlbumService
	trackService  TrackService
	creditService CreditService
}

func NewAlbumPolicy(
	albumService AlbumService,
	trackService TrackService,
	creditService CreditService) *albumPolicy {
	return &albumPolicy{
		albumService:  albumService,
		trackService:  trackService,
		creditService: creditService}
}

func (p *albumPolicy) GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumView, error) {
	products, err := p.albumService.GetAll(ctx, filtering, sorting)
	if err != nil {
		return nil, errors.Wrap(err, "albumService.All")
	}

	return products, nil
}

func (p *albumPolicy) Delete(ctx context.Context, id string) error {
	return p.albumService.Delete(ctx, id)
}

func (p *albumPolicy) Update(ctx context.Context, album model.AlbumView) error {
	return p.albumService.Update(ctx, album)
}

func (p *albumPolicy) Create(ctx context.Context, album model.Album) error {
	return p.albumService.Create(ctx, album)
}

func (p *albumPolicy) GetOne(ctx context.Context, albumID string) (model.FullAlbum, error) {
	album, err := p.albumService.GetOne(ctx, albumID)
	if err != nil {
		return model.FullAlbum{}, err
	}
	credits, err := p.creditService.GetAll(ctx, albumID)
	if err != nil {
		return model.FullAlbum{}, err
	}
	tracklist, err := p.trackService.GetAll(ctx, albumID)
	if err != nil {
		return model.FullAlbum{}, err
	}
	return model.FullAlbum{
		Album:     album.Album,
		Info:      album.Info,
		Credits:   credits,
		Tracklist: tracklist,
	}, nil
}
