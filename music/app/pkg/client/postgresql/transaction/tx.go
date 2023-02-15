package transaction

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Tx struct {
	pool *pgxpool.Pool
}

func (t *Tx) WithinTransaction(ctx context.Context, isoLevel pgx.TxOptions, tFunc func(ctx context.Context) error) error {
	tx, err := t.pool.BeginTx(ctx, isoLevel)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = tFunc(ctx)
	if err != nil {
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
