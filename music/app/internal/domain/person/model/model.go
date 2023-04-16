package model

import (
	"errors"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
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

func NewPersonFromPB(request *albumPb.CreatePersonRequest) Person {
	return Person{
		FirstName: request.GetFirstName(),
		LastName:  request.GetLastName(),
		BirthDate: request.GetBirthDate(),
	}
}

func (p Person) ToProto() *albumPb.Person {
	return &albumPb.Person{
		PersonId:  p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		BirthDate: p.BirthDate,
	}
}
