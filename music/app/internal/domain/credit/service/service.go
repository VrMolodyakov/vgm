package service

import (
	"context"
	"errors"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type CreditDAO interface {
	Create(ctx context.Context, credits []model.Credit) error
	GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error)
	Delete(ctx context.Context, albumID string) error
}

type creditService struct {
	creditDAO CreditDAO
}

func NewCreditService(dao CreditDAO) *creditService {
	return &creditService{creditDAO: dao}
}

func (c *creditService) Create(ctx context.Context, credits []model.Credit) error {
	if len(credits) == 0 {
		return errors.New("empty credits")
	}
	return c.creditDAO.Create(ctx, credits)
}

func (c *creditService) GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error) {
	if albumID == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return nil, err
	}
	informations, err := c.creditDAO.GetAll(ctx, albumID)
	if err != nil {
		return nil, err
	}
	return informations, nil
}

func (c *creditService) Delete(ctx context.Context, albumID string) error {
	if albumID == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err
	}
	return c.creditDAO.Delete(ctx, albumID)
}
