package user_service

import (
	"http2/app/types/userDB"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	SaveUser(val user_types.User) (*user_types.User, error)
	GetUser(val user_types.Credential) (*user_types.User, error)
	GetAllUser() ([]user_types.User, error)
	Update(id uint64, val map[string]interface{}) (*user_types.UpdateUserResponseDTO, error)
	Delete(id uint64) error
	GetUserByLogin(str string) (*user_types.User, error)
	GetUserByID(id uint64) (*user_types.User, error)
	GetUserByIDs(ids []int) ([]user_types.User, error)
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

