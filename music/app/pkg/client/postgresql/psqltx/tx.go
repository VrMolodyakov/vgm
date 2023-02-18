package psqltx

import (
	"context"
	"fmt"
	"log"

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

func (r *tx) Conn() (db db.PostgreSQLClient) {
	return r.db
}

func (t *tx) WithinTransaction(
	ctx context.Context,
	construct func(tx db.PostgreSQLClient) Transactor,
	action func(txRepo Transactor) error) (err error) {

	// return t.doTx(func(ctx context.Context, tx Transactor) error {
	// 	return action(construct(tx.Conn()))
	// })
	return t.doTx(ctx, func(ctx context.Context, tx Transactor) error {
		return action(construct(tx.Conn()))
	})
}

func (t *tx) doTx(ctx context.Context, fn func(context.Context, Transactor) error) error {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	fmt.Println("start of transaction")
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()
	err = fn(ctx, t.withTx(tx))
	fmt.Println("ERROR: ", err)
	if err != nil {
		fmt.Println("inside")
		if errRollback := tx.Rollback(ctx); errRollback != nil {
			log.Printf("rollback transaction: %v", errRollback)
		}
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (t *tx) withTx(db db.PostgreSQLClient) Transactor {
	return &tx{
		db: db,
	}
}
