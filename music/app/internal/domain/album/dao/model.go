package dao

import "time"

type AlbumStorage struct {
	ID       string
	Title    string
	CreateAt time.Time
}
