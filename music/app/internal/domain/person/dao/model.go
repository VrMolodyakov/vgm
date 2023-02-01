package dao

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/person/model"
	mapper "github.com/worldline-go/struct2"
)

type PersonStorage struct {
	ID        int64  `struct:"person_id "`
	FirstName string `struct:"first_name"`
	LastName  string `struct:"last_name"`
}

func fromModel(person model.Person) PersonStorage {
	return PersonStorage{}
}

func (p PersonStorage) ToModel() model.Person {
	return model.Person{}
}

func toStorageMap(person model.Person) map[string]interface{} {
	storage := fromModel(person)
	albumStorageMap := (&mapper.Decoder{}).Map(storage)
	return albumStorageMap
}
