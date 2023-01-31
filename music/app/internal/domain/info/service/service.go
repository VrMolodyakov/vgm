package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
)

type InfoDAO interface {
	GetAll(ctx context.Context) ([]dao.InfoStorage, error)
	Create(ctx context.Context, Info model.Info) (dao.InfoStorage, error)
	GetOne(ctx context.Context, albumID string) (dao.InfoStorage, error)
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

func (s *infoService) Create(ctx context.Context, album model.Info) (model.Info, error) {
	dbInfo, err := s.infoDAO.Create(ctx, album)
	if err != nil {
		return model.Info{}, errors.Wrap(err, "infoService.Create")
	}

	return dbInfo.ToModel(), nil
}

func (s *infoService) GetOne(ctx context.Context, albumID string) (model.Info, error) {
	dbInfo, err := s.infoDAO.GetOne(ctx, albumID)
	if err != nil {
		return model.Info{}, errors.Wrap(err, "infoService.Create")
	}

	return dbInfo.ToModel(), nil
}
