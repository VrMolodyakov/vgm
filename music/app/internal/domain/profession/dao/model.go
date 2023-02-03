package dao

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/profession/model"
	mapper "github.com/worldline-go/struct2"
)

type ProfessionStorage struct {
	ID    int64  `struct:"profession_id"`
	Title string `struct:"profession_title"`
}

func toStorageMap(proff model.Profession) map[string]interface{} {
	storage := FromModel(proff)
	professionStorageMap := (&mapper.Decoder{}).Map(storage)
	return professionStorageMap
}

func FromModel(m model.Profession) ProfessionStorage {
	return ProfessionStorage{
		ID:    m.ID,
		Title: m.Title,
	}
}

func (p ProfessionStorage) ToModel() model.Profession {
	return model.Profession{
		ID:    p.ID,
		Title: p.Title,
	}
}
