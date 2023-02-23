package reposotory

import (
	"time"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	mapper "github.com/worldline-go/struct2"
)

const (
	fields int = 5
)

type personStorage struct {
	ID        int64
	FirstName string
	LastName  string
	BirthDate time.Time
}

type queryStorage struct {
	FirstName string    `struct:"first_name"`
	LastName  string    `struct:"last_name"`
	BirthDate time.Time `struct:"birth_date"`
}

func toStorage(person model.Person) queryStorage {
	birthDate := time.UnixMilli(person.BirthDate)
	return queryStorage{
		FirstName: person.FirstName,
		LastName:  person.LastName,
		BirthDate: birthDate,
	}
}

func (p personStorage) toModel() model.Person {
	return model.Person{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		BirthDate: p.BirthDate.UnixMilli(),
	}
}

func toStorageMap(person model.Person) map[string]interface{} {
	storage := toStorage(person)
	personStorageMap := (&mapper.Decoder{}).Map(storage)
	return personStorageMap
}

func toUpdateStorageMap(m model.Person) map[string]interface{} {

	storageMap := make(map[string]interface{}, fields)

	if m.FirstName != "" {
		storageMap["first_name"] = m.FirstName
	}
	if m.LastName != "" {
		storageMap["last_name"] = m.LastName
	}
	if m.BirthDate != 0 {
		storageMap["birth_date"] = m.BirthDate
	}
	return storageMap
}
