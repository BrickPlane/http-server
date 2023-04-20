package wallet

import (
	"http2/app/storage"
	wallet_types "http2/app/types/walletDB"

	"github.com/jmoiron/sqlx"
)

type WalletStorage struct {
	DB *sqlx.DB
}

func NewWalletStorage() (*WalletStorage, error) {
	con, err := storage.ConnectDB()
	if err != nil {
		return nil, err 
	}
	return &WalletStorage{DB: con}, nil
} 

func (storage *WalletStorage) ReplenishmentWallet(val wallet_types.Replenishment) error {
	_, err := storage.DB.Query(`UPDATE "users" SET wallet = wallet + $1 WHERE id =$2;`,val.Fill, val.ID)

	if err != nil {
		return err
	}
	return nil
}