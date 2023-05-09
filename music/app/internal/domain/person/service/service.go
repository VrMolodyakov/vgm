package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("person-service")
)

type PersonRepo interface {
	GetAll(ctx context.Context, filtering filter.Filterable) ([]model.Person, error)
	Create(ctx context.Context, person model.Person) (model.Person, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, person model.Person) error
}

type personService struct {
	personRepo PersonRepo
}

func NewPersonService(dao PersonRepo) *personService {
	return &personService{personRepo: dao}
}

func (p *personService) GetAll(ctx context.Context, filter filter.Filterable) ([]model.Person, error) {
	ctx, span := tracer.Start(ctx, "service.GetAll")
	defer span.End()

	persons, err := p.personRepo.GetAll(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "personService.All")
	}
	return persons, nil

}

func (p *personService) Create(ctx context.Context, person model.Person) (model.Person, error) {
	ctx, span := tracer.Start(ctx, "service.Create")
	defer span.End()

	if !person.IsValid() {
		return model.Person{}, model.ErrValidation
	}
	person, err := p.personRepo.Create(ctx, person)
	if err != nil {
		return model.Person{}, errors.Wrap(err, "personService.Create")
	}

	return person, nil
}

func (p *personService) Update(ctx context.Context, person model.Person) error {
	ctx, span := tracer.Start(ctx, "service.Update")
	defer span.End()

	if !person.IsValid() {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err
	}
	err := p.personRepo.Update(ctx, person)
	if err != nil {
		return errors.Wrap(err, "trackService.Update")
	}
	return nil
}

func (p *personService) Delete(ctx context.Context, id string) error {
	ctx, span := tracer.Start(ctx, "service.Delete")
	defer span.End()

	if id == "" {
		err := errors.New("id must not be empty")
		logging.LoggerFromContext(ctx).Error(err.Error())
		return err

	}
	return p.personRepo.Delete(ctx, id)
}
