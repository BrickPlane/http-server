package purch_storage

import (
	"database/sql"
	"http2/app/storage"
	"http2/app/types/productDB"
	"http2/app/types/purchases"
	"http2/app/types/userDB"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	DB *sqlx.DB
}

func NewStorage() (*Storage, error) {
	con, err := storage.ConnectDB()
	if err != nil {
		return nil, err
	}

	return &Storage{DB: con}, nil
}

func (stor *Storage) Receipt(user *user_types.User, product *product_types.ProductsResponce, tx *sql.Tx) error {
	q := `
	INSERT INTO "purchases" (id_buyer, id_goods, money_spended)	
	VALUES ($1, $2, $3);
	`
	_, err := tx.Exec(q, user.ID, product.ID, product.Price)
	if err != nil {
		tx.Rollback()
		return err
		
	}

	tx.Commit()
	return nil
}

func (stor *Storage) GetPurchased() ([]purchases_type.Purchases, error) {
	query, err := stor.DB.Query(`SELECT * FROM "purchases"`)
	if err != nil {
		return nil, err
	}

	slice := make([]purchases_type.Purchases, 0)
	for query.Next() {
		var data purchases_type.Purchases
		if err := query.Scan(
			&data.IdBuyer,
			&data.IdGoods,
			&data.MoneySpended,
		); err != nil {
			return nil, err
		}
		slice = append(slice, data)
	}

	return slice, nil
}
