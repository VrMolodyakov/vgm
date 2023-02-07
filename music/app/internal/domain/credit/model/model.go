package model

import (
	"errors"

	creditPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/credit/v1"
)

var (
	ErrValidation = errors.New("ID's must not be null")
)

type Credit struct {
	PersonID     int64
	AlbumID      string
	ProfessionID int64
}

type CreditInfo struct {
	FirstName  string
	LastName   string
	Profession string
}

func (c *Credit) IsEmpty() bool {
	return c.AlbumID == "" || c.PersonID == 0 || c.ProfessionID == 0
}

func NewCreditFromPB(pb *creditPb.CreateCreditRequest) Credit {
	return Credit{
		PersonID: pb.GetPersonId(),
		AlbumID:  pb.GetAlbumId(),
	}
}
