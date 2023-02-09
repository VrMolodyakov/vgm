package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/info/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type InfoDAO interface {
	Create(ctx context.Context, Info model.Info) (model.Info, error)
	GetOne(ctx context.Context, albumID string) (model.Info, error)
	Update(ctx context.Context, info model.Info) error
	Delete(ctx context.Context, id string) error
}

type infoService struct {
	infoDAO InfoDAO
}

func NewInfoService(dao InfoDAO) *infoService {
	return &infoService{infoDAO: dao}
}

func (s *infoService) Create(ctx context.Context, info model.Info) (model.Info, error) {
	if info.IsEmpty() {
		return model.Info{}, model.ErrValidation
	}
	info, err := s.infoDAO.Create(ctx, info)
	if err != nil {
		return model.Info{}, errors.Wrap(err, "infoService.Create")
	}

	return info, nil
}

func (s *infoService) GetOne(ctx context.Context, infoID string) (model.Info, error) {
	if infoID == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return model.Info{}, err
	}
	info, err := s.infoDAO.GetOne(ctx, infoID)
	if err != nil {
		return model.Info{}, errors.Wrap(err, "infoService.Create")
	}

	return info, nil
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
