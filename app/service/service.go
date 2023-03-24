package service

import (
	"http2/app/types"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	SaveUser(val types.Credential) (*types.Credential, error)
	GetUser() ([]types.Credential, error)
	Update(val types.Credential) (*types.Credential, error)
	Delete(val types.Credential) error
	GetUserByID(val types.Credential) (*types.Credential, error)
	GetUserByIDs(ids []int) ([]types.Credential, error)
}

type Service struct {
	DB      *sqlx.DB
	storage IStorage
}

func NewService(storage IStorage, conn *sqlx.DB) *Service {
	return &Service{
		storage: storage,
		DB:      conn,
	}
}
