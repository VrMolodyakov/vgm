package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type InfoDAO interface {
	GetAll(ctx context.Context) ([]dao.InfoStorage, error)
	Create(ctx context.Context, Info model.Info) (dao.InfoStorage, error)
	GetOne(ctx context.Context, albumID string) (dao.InfoStorage, error)
	Update(ctx context.Context, info model.Info) error
	Delete(ctx context.Context, id string) error
}

type infoService struct {
	infoDAO InfoDAO
}

func NewInfoService(dao InfoDAO) *infoService {
	return &infoService{infoDAO: dao}
}

func (a *infoService) GetAll(ctx context.Context) ([]*model.Info, error) {
	dbAlbumsInfo, err := a.infoDAO.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "infoService.All")
	}
	var albumsInfo []*model.Info
	for _, info := range dbAlbumsInfo {
		m := info.ToModel()
		albumsInfo = append(albumsInfo, &m)
	}
	return albumsInfo, nil

}

func (s *infoService) Create(ctx context.Context, info model.Info) (model.Info, error) {
	if info.IsEmpty() {
		return model.Info{}, model.ErrValidation
	}
	dbInfo, err := s.infoDAO.Create(ctx, info)
	if err != nil {
		return model.Info{}, errors.Wrap(err, "infoService.Create")
	}

	return dbInfo.ToModel(), nil
}

func (s *infoService) GetOne(ctx context.Context, infoID string) (model.Info, error) {
	if infoID == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return model.Info{}, err
	}
	dbInfo, err := s.infoDAO.GetOne(ctx, infoID)
	if err != nil {
		return model.Info{}, errors.Wrap(err, "infoService.Create")
	}

	return dbInfo.ToModel(), nil
}

func (s *infoService) Update(ctx context.Context, info model.Info) error {
	if info.IsEmpty() {
		return model.ErrValidation
	}
	return s.infoDAO.Update(ctx, info)
}

func (s *infoService) Delete(ctx context.Context, id string) error {
	if id == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err
	}

	return s.infoDAO.Delete(ctx, id)
}
