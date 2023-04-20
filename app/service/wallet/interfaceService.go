package wallet_service

import (
	"http2/app/types/walletDB"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	ReplenishmentWallet(val wallet_types.Replenishment) error
}

type Service struct {
	DB *sqlx.DB
	storage IStorage
}

func NewService(storage IStorage, conn *sqlx.DB) *Service{
	return &Service{
		storage: storage,
		DB: conn,
	}
}