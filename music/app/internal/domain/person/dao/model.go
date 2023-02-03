package dao

import (
	"time"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	mapper "github.com/worldline-go/struct2"
)

type PersonStorage struct {
	ID        int64     `struct:"person_id "`
	FirstName string    `struct:"first_name"`
	LastName  string    `struct:"last_name"`
	BirthDate time.Time `struct:"birth_date"`
}

func fromModel(person model.Person) PersonStorage {
	birthDate := time.UnixMilli(person.BirthDate)
	return PersonStorage{
		ID:        person.ID,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		BirthDate: birthDate,
	}
}

func (p PersonStorage) ToModel() model.Person {
	return model.Person{}
}

func toStorageMap(person model.Person) map[string]interface{} {
	storage := fromModel(person)
	personStorageMap := (&mapper.Decoder{}).Map(storage)
	return personStorageMap
}
