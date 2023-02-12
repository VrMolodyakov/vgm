package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type AlbumDAO interface {
	GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.AlbumView, error)
	GetOne(ctx context.Context, albumID string) (model.AlbumView, error)
	Create(ctx context.Context, album model.AlbumView) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, album model.AlbumView) error
}

type albumService struct {
	albumDAO AlbumDAO
}

func NewAlbumService(dao AlbumDAO) *albumService {
	return &albumService{albumDAO: dao}
}

func (a *albumService) GetAll(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.AlbumView, error) {
	albums, err := a.albumDAO.GetAll(ctx, filter, sort)
	if err != nil {
		return nil, errors.Wrap(err, "albumService.All")
	}
	return albums, nil

}

func (s *albumService) Create(ctx context.Context, album model.AlbumView) error {
	if album.IsEmpty() {
		return model.ErrValidation
	}
	return s.albumDAO.Create(ctx, album)
}

func (s *albumService) Delete(ctx context.Context, id string) error {
	if id == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err

	}
	return s.albumDAO.Delete(ctx, id)
}

func (s *albumService) Update(ctx context.Context, album model.AlbumView) error {
	if album.IsEmpty() {
		return model.ErrValidation
	}
	return s.albumDAO.Update(ctx, album)
}

func (s *albumService) GetOne(ctx context.Context, albumID string) (model.AlbumView, error) {
	if albumID == "" {
		return model.AlbumView{}, errors.New("album id is empty")
	}
	return s.albumDAO.GetOne(ctx, albumID)

}
