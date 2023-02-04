package model

import "errors"

var (
	ErrValidation = errors.New("title must not be empty")
)

type Track struct {
	ID      int64
	AlbumID string
	Title   string
}

func (t *Track) IsEmpty() bool {
	return t.Title == ""
}
