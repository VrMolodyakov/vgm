package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/profession/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
)

type ProfessionDAO interface {
	Create(ctx context.Context, profession string) (model.Profession, error)
	GetOne(ctx context.Context, profID string) (model.Profession, error)
}

type professionService struct {
	professionDAO ProfessionDAO
}

func NewProfessionService(dao ProfessionDAO) *professionService {
	return &professionService{professionDAO: dao}
}

func (s *professionService) Create(ctx context.Context, prof string) (model.Profession, error) {
	if prof == "" {
		return model.Profession{}, errors.New("profession id must not be null")
	}
	profession, err := s.professionDAO.Create(ctx, prof)
	if err != nil {
		return model.Profession{}, errors.Wrap(err, "ProfessionService.Create")
	}

	return profession, nil
}

func (s *professionService) GetOne(ctx context.Context, profID string) (model.Profession, error) {
	if profID == "" {
		return model.Profession{}, errors.New("profession id must not be null")
	}
	profession, err := s.professionDAO.GetOne(ctx, profID)
	if err != nil {
		return model.Profession{}, errors.Wrap(err, "ProfessionService.GetOne")
	}

	return profession, nil
}
