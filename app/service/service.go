package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IStorage interface {
	StorageIn(c *gin.Context, data string) error
}

type Service struct{
	storage IStorage
}

func NewService(storage IStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) SignToken(c *gin.Context) string {
	x := "magic"
	err := s.storage.StorageIn(c, x)
	if err != nil {
		return fmt.Sprintln("err")
	}
	return x
}
