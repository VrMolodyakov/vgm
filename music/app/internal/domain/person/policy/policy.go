package policy

import (
	"context"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	"github.com/VrMolodyakov/vgm/music/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("person-policy")
)

type PersonService interface {
	GetAll(ctx context.Context, filter filter.Filterable) ([]model.Person, error)
	Create(ctx context.Context, Person model.Person) (model.Person, error)
	Update(ctx context.Context, person model.Person) error
	Delete(ctx context.Context, id string) error
}

type PersonPolicy struct {
	personService PersonService
}

func NewPersonPolicy(service PersonService) *PersonPolicy {
	return &PersonPolicy{personService: service}
}

func (p *PersonPolicy) GetAll(ctx context.Context, filter filter.Filterable) ([]model.Person, error) {
	ctx, span := tracer.Start(ctx, "policy.GetAll")
	defer span.End()
	products, err := p.personService.GetAll(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "PersonService.All")
	}

	return products, nil
}

func (p *PersonPolicy) Create(ctx context.Context, person model.Person) (model.Person, error) {
	ctx, span := tracer.Start(ctx, "policy.Create")
	defer span.End()
	return p.personService.Create(ctx, person)
}

func (p *PersonPolicy) Update(ctx context.Context, person model.Person) error {
	ctx, span := tracer.Start(ctx, "policy.Update")
	defer span.End()
	return p.personService.Update(ctx, person)
}

func (p *PersonPolicy) Delete(ctx context.Context, albumID string) error {
	ctx, span := tracer.Start(ctx, "policy.Delete")
	defer span.End()
	return p.personService.Delete(ctx, albumID)
}
