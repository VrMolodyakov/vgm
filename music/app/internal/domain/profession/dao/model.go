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
	ProfessionStorageMap := (&mapper.Decoder{}).Map(storage)
	return ProfessionStorageMap
}

func FromModel(m model.Profession) ProfessionStorage {
	return ProfessionStorage{
		ID:    m.ID,
		Title: m.Title,
	}
}
