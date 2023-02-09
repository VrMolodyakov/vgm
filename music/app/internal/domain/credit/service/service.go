package service

import (
	"context"
	"errors"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type CreditDAO interface {
	Create(ctx context.Context, credit model.Credit) (model.Credit, error)
	GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error)
	Delete(ctx context.Context, albumID string) error
}

type creditService struct {
	creditDAO CreditDAO
}

func NewCreditService(dao CreditDAO) *creditService {
	return &creditService{creditDAO: dao}
}

func (c *creditService) Create(ctx context.Context, credit model.Credit) (model.Credit, error) {
	if credit.IsEmpty() {
		return model.Credit{}, model.ErrValidation
	}
	info, err := c.creditDAO.Create(ctx, credit)
	if err != nil {
		return model.Credit{}, err
	}
	return info, nil

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