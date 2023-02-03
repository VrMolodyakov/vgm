package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/profession/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/profession/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
)

type ProfessionDAO interface {
	Create(ctx context.Context, profession model.Profession) (dao.ProfessionStorage, error)
	GetOne(ctx context.Context, profID string) (dao.ProfessionStorage, error)
}

type professionService struct {
	professionDAO ProfessionDAO
}

func NewProfessionService(dao ProfessionDAO) *professionService {
	return &professionService{professionDAO: dao}
}

func (s *professionService) Create(ctx context.Context, profession model.Profession) (model.Profession, error) {
	if profession.IsEmpty() {
		return model.Profession{}, model.ErrValidation
	}
	dbProfession, err := s.professionDAO.Create(ctx, profession)
	if err != nil {
		return model.Profession{}, errors.Wrap(err, "ProfessionService.Create")
	}

	return dbProfession.ToModel(), nil
}

func (s *professionService) GetOne(ctx context.Context, profID string) (model.Profession, error) {
	if profID == "" {
		return model.Profession{}, errors.New("profession id must not be null")
	}
	dbProfession, err := s.professionDAO.GetOne(ctx, profID)
	if err != nil {
		return model.Profession{}, errors.Wrap(err, "ProfessionService.GetOne")
	}

	return dbProfession.ToModel(), nil
}
