package purch_storage

import (
	"context"
	"database/sql"
)

func (stor *Storage) Transaction() (*sql.Tx, error) {
	ctx := context.Background()
	tx, err := stor.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return tx, nil
}

func (stor *Storage) ChangeWallet(set float64, id int, tx *sql.Tx) error {
	q := `UPDATE "users" SET wallet=$1 WHERE id=$2`
	_, err := tx.Exec(q, set, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
