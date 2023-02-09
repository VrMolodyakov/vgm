package model

import "errors"

var (
	ErrValidation = errors.New("title must not be empty")
)

type Profession struct {
	ID    int64
	Title string
}

func (p *Profession) IsEmpty() bool {
	return p.Title == ""
}
