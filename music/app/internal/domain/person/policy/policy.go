package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

type PersonService interface {
	GetAll(ctx context.Context, filter filter.Filterable, sort sort.Sortable) ([]model.Person, error)
	Create(ctx context.Context, Person model.Person) (model.Person, error)
}

type PersonPolicy struct {
	personService PersonService
}

func NewPersonPolicy(service PersonService) *PersonPolicy {
	return &PersonPolicy{personService: service}
}

func (p *PersonPolicy) GetAll(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.Person, error) {
	products, err := p.personService.GetAll(ctx, filtering, sorting)
	if err != nil {
		return nil, errors.Wrap(err, "PersonService.All")
	}

	return products, nil
}

func (p *PersonPolicy) Create(ctx context.Context, person model.Person) (model.Person, error) {
	return p.personService.Create(ctx, person)
}
