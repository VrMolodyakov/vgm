package psqltx

import (
	"context"
	"fmt"

	db "github.com/VrMolodyakov/vgm/music/app/pkg/client/postgresql"
)

type tx struct {
	db db.PostgreSQLClient
}

type Transactor interface {
	WithinTransaction(
		ctx context.Context,
		construct func(tx db.PostgreSQLClient) Transactor,
		action func(txRepo Transactor) error) (err error)
	Conn() (db db.PostgreSQLClient)
}

func NewTx(db db.PostgreSQLClient) *tx {
	return &tx{
		db: db,
	}
}

func (r *tx) Conn() db.PostgreSQLClient {
	return r.db
}

func (t *tx) WithinTransaction(
	ctx context.Context,
	construct func(tx db.PostgreSQLClient) Transactor,
	action func(txRepo Transactor) error) (err error) {

	return t.doTx(ctx, func(tx Transactor) error {
		return action(construct(tx.Conn()))
	})
}

func (t *tx) doTx(ctx context.Context, fn func(Transactor) error) error {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			fmt.Println("-----ROLLBACK-----")
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	err = fn(t.withTx(tx))
	if err != nil {
		fmt.Println("-----GOT ERROR-----")
		return err
	}
	return nil
}

func (t *tx) withTx(db db.PostgreSQLClient) Transactor {
	return &tx{
		db: db,
	}
}
