package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("album-service")
)

type albumService struct {
	albumRepo AlbumRepo
}

func NewAlbumService(albumRepo AlbumRepo) *albumService {
	return &albumService{albumRepo: albumRepo}
}

func (a *albumService) GetAll(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.AlbumPreview, error) {
	ctx, span := tracer.Start(ctx, "service.GetAll")
	defer span.End()

	albums, err := a.albumRepo.GetAll(ctx, filter, sort)
	if err != nil {
		return nil, errors.Wrap(err, "albumService.All")
	}
	return albums, nil

}

func (s *albumService) Delete(ctx context.Context, id string) error {
	ctx, span := tracer.Start(ctx, "service.Delete")
	defer span.End()

	if id == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err

	}
	return s.albumRepo.Delete(ctx, id)
}

func (s *albumService) Update(ctx context.Context, album model.AlbumView) error {
	ctx, span := tracer.Start(ctx, "service.Update")
	defer span.End()

	if album.IsEmpty() {
		return model.ErrValidation
	}
	return s.albumRepo.Update(ctx, album)
}

func (s *albumService) GetOne(ctx context.Context, albumID string) (model.AlbumInfo, error) {
	ctx, span := tracer.Start(ctx, "service.GetOne")
	defer span.End()

	if albumID == "" {
		return model.AlbumInfo{}, errors.New("album id is empty")
	}
	return s.albumRepo.GetInfo(ctx, albumID)

}

func (s *albumService) Create(ctx context.Context, album model.Album) error {
	ctx, span := tracer.Start(ctx, "service.Create")
	defer span.End()

	return s.albumRepo.Create(ctx, album)
}

func (s *albumService) GetDays(ctx context.Context, count uint64) ([]int64, error) {
	ctx, span := tracer.Start(ctx, "service.GetDays")
	defer span.End()
	dates, err := s.albumRepo.GetLastDays(ctx, count)
	if err != nil {
		return nil, err
	}
	unix := make([]int64, len(dates))
	for i := range dates {
		unix[i] = dates[i].UnixMilli()
	}
	return unix, nil
}
