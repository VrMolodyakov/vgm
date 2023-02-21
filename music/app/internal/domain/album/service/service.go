package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type albumService struct {
	albumRepo  AlbumRepo
	creditRepo CreditRepo
	infoRepo   InfoRepo
	trackRepo  TrackRepo
}

func NewAlbumService(
	albumRepo AlbumRepo,
	creditRepo CreditRepo,
	infoRepo InfoRepo,
	trackRepo TrackRepo) *albumService {
	return &albumService{albumRepo: albumRepo, creditRepo: creditRepo, infoRepo: infoRepo, trackRepo: trackRepo}
}

func (a *albumService) GetAll(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.AlbumView, error) {
	albums, err := a.albumRepo.GetAll(ctx, filter, sort)
	if err != nil {
		return nil, errors.Wrap(err, "albumService.All")
	}
	return albums, nil

}

func (s *albumService) Delete(ctx context.Context, id string) error {
	if id == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err

	}
	return s.albumRepo.Delete(ctx, id)
}

func (s *albumService) Update(ctx context.Context, album model.AlbumView) error {
	if album.IsEmpty() {
		return model.ErrValidation
	}
	return s.albumRepo.Update(ctx, album)
}

func (s *albumService) GetOne(ctx context.Context, albumID string) (model.AlbumView, error) {
	if albumID == "" {
		return model.AlbumView{}, errors.New("album id is empty")
	}
	return s.albumRepo.GetOne(ctx, albumID)

}

func (s *albumService) Create(ctx context.Context, album model.Album) error {
	return s.albumRepo.Create(ctx, album)
}
