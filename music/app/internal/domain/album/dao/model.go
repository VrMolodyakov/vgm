package dao

import "time"

type AlbumStorage struct {
	ID         string
	Title      string
	ReleasedAt time.Time
	CreatedAt  time.Time
}
