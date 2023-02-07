package service

import (
	"context"
	"errors"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type CreditDAO interface {
	Create(ctx context.Context, credit model.Credit) (dao.CreditStorage, error)
	GetAll(ctx context.Context, albumID string) ([]dao.CreditInfoStorage, error)
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
	srorage, err := c.creditDAO.Create(ctx, credit)
	if err != nil {
		return model.Credit{}, err
	}
	return srorage.ToModel(), nil

}

func (c *creditService) GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error) {
	if albumID == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return nil, err
	}
	storageList, err := c.creditDAO.GetAll(ctx, albumID)
	if err != nil {
		return nil, err
	}
	informations := make([]model.CreditInfo, len(storageList))
	for i := 0; i < len(storageList); i++ {
		informations[i] = storageList[i].ToModel()
	}
	return informations, nil
}
