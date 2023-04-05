package controller

import (
	"http2/app/types"

	"github.com/gin-gonic/gin"
)


type IService interface {
	SigninUser(creds types.User) (*types.User, error)
	Login(creds types.Credential) (*types.User, error)
	GenToken(c *gin.Context, creds types.User) (string, error)
	ParseWithBearer(c *gin.Context) (*types.JWTUploadData, error)
	TokenVerification(tokenData *types.JWTUploadData) error
	GetAllUser() ([]types.User, error)
	GetUserByLogin(str string) (*types.User, error)
	UpdateUser(upd types.User) (*types.User, error)
	DeleteUser(dlt uint64) error
	GetUserByID(id uint64) (*types.User, error)
	GetUserByIDs(ids []int) ([]types.User, error)
}

type Controller struct {
	service IService

}

func NewController(service IService) *Controller {
	return &Controller{
		service: service,
	}
}
