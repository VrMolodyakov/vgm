package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type repository interface {
	All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]dao.AlbumStorage, error)
	Create(ctx context.Context, m map[string]interface{}) (dao.AlbumStorage, error)
}

type albumService struct {
	repository repository
}

func NewAlbumService(repository repository) *albumService {
	return &albumService{repository: repository}
}

func (a *albumService) All(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.Album, error) {
	dbAlbums, err := a.repository.All(ctx, filter, sort)
	if err != nil {
		return nil, errors.Wrap(err, "albumService.All")
	}
	var albums []model.Album
	for _, album := range dbAlbums {
		albums = append(albums, model.NewAlbum(album))
	}
	return albums, nil

}

func (s *albumService) Create(ctx context.Context, album model.Album) (model.Album, error) {
	albumStorageMap, err := album.ToMap()
	if err != nil {
		return model.Album{}, err
	}

	dbAlbum, err := s.repository.Create(ctx, albumStorageMap)
	if err != nil {
		return model.Album{}, errors.Wrap(err, "albumService.Create")
	}

	return model.NewAlbum(dbAlbum), nil
}
