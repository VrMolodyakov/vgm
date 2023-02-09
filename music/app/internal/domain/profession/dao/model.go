package dao

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/profession/model"
	mapper "github.com/worldline-go/struct2"
)

type professionStorage struct {
	ID    int64  `struct:"profession_id"`
	Title string `struct:"profession_title"`
}

func toStorageMap(proff model.Profession) map[string]interface{} {
	storage := fromModel(proff)
	professionStorageMap := (&mapper.Decoder{}).Map(storage)
	return professionStorageMap
}

func fromModel(m model.Profession) professionStorage {
	return professionStorage{
		ID:    m.ID,
		Title: m.Title,
	}
}

func (p professionStorage) toModel() model.Profession {
	return model.Profession{
		ID:    p.ID,
		Title: p.Title,
	}
}
