package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type repository interface {
	All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]dao.AlbumStorage, error)
	Create(ctx context.Context, m map[string]interface{}) error
}

type albumService struct {
	repository repository
}

func NewAlbumService(repository repository) *albumService {
	return &albumService{repository: repository}
}

func (a *albumService) All(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.Album, error) {
	// dbAlbums ,err := a.repository.All(ctx,filter,sort)
	// if err != nil {
	// 	return nil, errors.Wrap(err,"albumService.All")
	// }
	// var albums []model.Album
	// for _,album := range dbAlbums{
	// albums = append(albums, album)
	// }
	return nil, nil

}
