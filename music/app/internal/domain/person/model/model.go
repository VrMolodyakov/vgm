package model

import (
	"errors"

	personPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/person/v1"
)

type Person struct {
	ID        int64
	FirstName string
	LastName  string
	BirthDate int64
}

var (
	ErrValidation = errors.New("first or last name must not be empty")
)

func (p *Person) IsValid() bool {
	return p.LastName != "" && p.FirstName != ""
}

func NewAlbumFromPB(pb *personPb.CreatePersonRequest) Person {
	return Person{
		FirstName: pb.GetFirstName(),
		LastName:  pb.GetLastName(),
		BirthDate: pb.GetBirthDate(),
	}
}

func (p Person) ToProto() *personPb.Person {
	return &personPb.Person{
		PersonId:  p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		BirthDate: p.BirthDate,
	}
}
