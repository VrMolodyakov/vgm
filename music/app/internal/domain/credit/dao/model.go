package dao

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
	mapper "github.com/worldline-go/struct2"
)

type CreditStorage struct {
	PersonID     int64  `struct:"person_id"`
	AlbumID      string `struct:"album_id"`
	ProfessionID int64  `struct:"profession_id"`
}

func fromModel(person model.Credit) CreditStorage {
	return CreditStorage{
		PersonID:     person.PersonID,
		AlbumID:      person.AlbumID,
		ProfessionID: person.ProfessionID,
	}
}

func (c CreditStorage) ToModel() model.Credit {
	return model.Credit{
		PersonID:     c.PersonID,
		AlbumID:      c.AlbumID,
		ProfessionID: c.ProfessionID,
	}
}

func toStorageMap(credit model.Credit) map[string]interface{} {
	storage := fromModel(credit)
	creditStorageMap := (&mapper.Decoder{}).Map(storage)
	return creditStorageMap
}
