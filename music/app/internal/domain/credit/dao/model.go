package dao

import (
	"github.com/VrMolodyakov/vgm/music/app/internal/domain/credit/model"
)

type CreditStorage struct {
	PersonID int64  `struct:"person_id"`
	AlbumID  string `struct:"album_id"`
	Role     string `struct:"credit_role"`
}

type CreditInfoStorage struct {
	Profession string
	FirstName  string
	LastName   string
}

func (c CreditInfoStorage) toModel() model.CreditInfo {
	return model.CreditInfo{
		Profession: c.Profession,
		LastName:   c.LastName,
		FirstName:  c.FirstName,
	}
}
