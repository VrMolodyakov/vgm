package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

type CreditRepo interface {
	GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error)
	Delete(ctx context.Context, albumID string) error
	Update(ctx context.Context, albumId string, role string) error
}

type creditService struct {
	creditRepo CreditRepo
}

func NewCreditService(dao CreditRepo) *creditService {
	return &creditService{creditRepo: dao}
}

func (c *creditService) GetAll(ctx context.Context, albumID string) ([]model.CreditInfo, error) {
	if albumID == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return nil, err
	}
	informations, err := c.creditRepo.GetAll(ctx, albumID)
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
	return c.creditRepo.Delete(ctx, albumID)
}

func (c *creditService) Update(ctx context.Context, albumId string, role string) error {
	if albumId == "" || role == "" {
		err := errors.New("id/role must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err
	}
	err := c.creditRepo.Update(ctx, albumId, role)
	if err != nil {
		return errors.Wrap(err, "trackService.Update")
	}
	return nil
}
