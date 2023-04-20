package purchases_service

import (
	"database/sql"
	"http2/app/types/productDB"
	"http2/app/types/purchases"
	"http2/app/types/userDB"

	"github.com/jmoiron/sqlx"
)

type IStorPurch interface {
	Receipt(user *user_types.User, product *product_types.ProductsResponce, tx *sql.Tx) error
	GetPurchased() ([]purchases_type.Purchases, error)
	Transaction() (*sql.Tx, error)
	ChangeWallet(set float64, id int, tx *sql.Tx) error
}

type IStorage interface {
	GetUserByID(id uint64) (*user_types.User, error)
}

type IStorProd interface {
	GetProductByID(id uint64) (*product_types.ProductsResponce, error)
}

type PurchService struct {
	DB           *sqlx.DB
	storagePurch IStorPurch
	storageUser IStorage
	storageProduct IStorProd
}

func NewPurchService(storage IStorPurch, stor IStorage, str IStorProd, con *sqlx.DB) *PurchService {
	return &PurchService{
		DB:           con,
		storagePurch: storage,
		storageUser: stor,
		storageProduct: str,

	}
}

// have (*purch_storage.Storage, *user_storage.Storage, *product_storage.ProdStorage, *sqlx.DB)
