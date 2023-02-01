package dao

import (
	sq "github.com/Masterminds/squirrel"
	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
)

type PersonDAO struct {
	queryBuilder sq.StatementBuilderType
	client       db.PostgreSQLClient
}

const (
	table = "person"
)

func NewPersonStorage(client db.PostgreSQLClient) *PersonDAO {
	return &PersonDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}
