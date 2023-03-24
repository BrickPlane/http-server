package controller

import (
	"http2/app/types"

	"github.com/gin-gonic/gin"
)


type IService interface {
	SigninUser(creds types.Credential) (*types.Credential, error)
	ParseWithBearer(c *gin.Context) error
	GetUser() ([]types.Credential, error)
	UpdateUser(upd types.Credential) (*types.Credential, error)
	DeleteUser(dlt types.Credential) error
	GetUserByID(get types.Credential) (*types.Credential, error)
	GetUserByIDs(gets []int) ([]types.Credential, error)
}

type Controller struct {
	service IService
}

func NewController(service IService) *Controller {
	return &Controller{
		service: service,
	}
}
