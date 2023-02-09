package dao

import (
	"time"

	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	mapper "github.com/worldline-go/struct2"
)

type PersonStorage struct {
	ID        int64
	FirstName string
	LastName  string
	BirthDate time.Time
}

type storage struct {
	FirstName string    `struct:"first_name"`
	LastName  string    `struct:"last_name"`
	BirthDate time.Time `struct:"birth_date"`
}

func toStorage(person model.Person) storage {
	birthDate := time.UnixMilli(person.BirthDate)
	return storage{
		FirstName: person.FirstName,
		LastName:  person.LastName,
		BirthDate: birthDate,
	}
}

// func fromModel(person model.Person) PersonStorage {
// 	birthDate := time.UnixMilli(person.BirthDate)
// 	return PersonStorage{
// 		ID:        person.ID,
// 		FirstName: person.FirstName,
// 		LastName:  person.LastName,
// 		BirthDate: birthDate,
// 	}
// }

func (p PersonStorage) ToModel() model.Person {
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
