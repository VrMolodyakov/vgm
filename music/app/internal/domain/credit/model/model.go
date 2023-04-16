package model

import (
	"errors"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
)

var (
	ErrValidation = errors.New("ID's must not be null")
)

type Credit struct {
	PersonID   int64
	AlbumID    string
	Profession string
}

type CreditInfo struct {
	FirstName  string
	LastName   string
	Profession string
}

func (c *CreditInfo) ToProto() albumPb.CreditInfo {
	return albumPb.CreditInfo{
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Profession: c.Profession,
	}
}

func (c *Credit) IsEmpty() bool {
	return c.AlbumID == "" || c.PersonID == 0 || c.Profession == ""
}

func NewCreditFromPB(pb *albumPb.Credit) Credit {
	return Credit{
		PersonID:   pb.GetPersonId(),
		Profession: pb.GetProfession(),
	}
}
