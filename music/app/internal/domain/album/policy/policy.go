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
	GetAll(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.AlbumView, error)
	GetOne(ctx context.Context, albumID string) (model.AlbumView, error)
	Create(ctx context.Context, album model.AlbumView) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.AlbumView) error
}

type InfoService interface {
	Create(ctx context.Context, info infoModel.Info) error
	GetOne(ctx context.Context, albumID string) (infoModel.Info, error)
}

type TrackService interface {
	Create(ctx context.Context, tracklist []trackModel.Track) error
	GetAll(ctx context.Context, albumID string) ([]trackModel.Track, error)
}

type CreditService interface {
	Create(ctx context.Context, credits []creditModel.Credit) error
	GetAll(ctx context.Context, albumID string) ([]creditModel.CreditInfo, error)
}

type albumPolicy struct {
	albumService  AlbumService
	infoService   InfoService
	trackService  TrackService
	creditService CreditService
}

func NewAlbumPolicy(
	albumService AlbumService,
	infoService InfoService,
	trackService TrackService,
	creditService CreditService) *albumPolicy {
	return &albumPolicy{
		albumService:  albumService,
		infoService:   infoService,
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
	err := p.albumService.Create(ctx, album.Album)
	if err != nil {
		return err
	}
	err = p.infoService.Create(ctx, album.Info)
	if err != nil {
		return err
	}

	err = p.trackService.Create(ctx, album.Tracklist)
	if err != nil {
		return err
	}

	err = p.creditService.Create(ctx, album.Credits)
	if err != nil {
		return err
	}

	return nil
}

func (p *albumPolicy) GetOne(ctx context.Context, albumID string) (model.FullAlbum, error) {
	album, err := p.albumService.GetOne(ctx, albumID)
	if err != nil {
		return model.FullAlbum{}, err
	}
	info, err := p.infoService.GetOne(ctx, albumID)
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
		Album:     album,
		Info:      info,
		Credits:   credits,
		Tracklist: tracklist,
	}, nil
}
