package service

import (
	"http2/app/types"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	SaveUser(val types.User) (*types.User, error)
	GetUser(val types.Credential) (*types.User, error)
	GetAllUser() ([]types.User, error)
	Update(id uint64, val map[string]interface{}) (*types.UpdateUserResponseDTO, error)
	Delete(id uint64) error
	GetUserByLogin(str string) (*types.User, error)
	GetUserByID(id uint64) (*types.User, error)
	GetUserByIDs(ids []int) ([]types.User, error)
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

