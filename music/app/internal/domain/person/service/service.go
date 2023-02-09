package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/dao"
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
)

type PersonDAO interface {
	GetAll(ctx context.Context, filtering filter.Filterable) ([]dao.PersonStorage, error)
	Create(ctx context.Context, person model.Person) (dao.PersonStorage, error)
}

type personService struct {
	personDAO PersonDAO
}

func NewPersonService(dao PersonDAO) *personService {
	return &personService{personDAO: dao}
}

func (p *personService) GetAll(ctx context.Context, filter filter.Filterable) ([]model.Person, error) {
	dbpersons, err := p.personDAO.GetAll(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "personService.All")
	}
	var persons []model.Person
	for _, person := range dbpersons {
		persons = append(persons, person.ToModel())
	}
	return persons, nil

}

func (p *personService) Create(ctx context.Context, person model.Person) (model.Person, error) {
	if person.IsEmpty() {
		return model.Person{}, model.ErrValidation
	}
	dbperson, err := p.personDAO.Create(ctx, person)
	if err != nil {
		return model.Person{}, errors.Wrap(err, "personService.Create")
	}

	return dbperson.ToModel(), nil
}
