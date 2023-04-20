package user_controller

import (
	"http2/app/types/userDB"
	"http2/app/types"

	"github.com/gin-gonic/gin"
)


type IService interface {
	SigninUser(creds user_types.User) (*user_types.User, error)
	Login(creds user_types.Credential) (*user_types.User, error)
	GenToken(c *gin.Context, creds user_types.User) (string, error)
	ParseWithBearer(c *gin.Context) (*types.JWTUploadData, error)
	TokenVerification(tokenData *types.JWTUploadData) error
	GetAllUser() ([]user_types.User, error)
	GetUserByLogin(str string) (*user_types.User, error)
	UpdateUser(upd user_types.UpdateUserRequestDTO) (*user_types.UpdateUserResponseDTO, error)
	DeleteUser(dlt uint64) error
	GetUserByID(id uint64) (*user_types.User, error)
	GetUserByIDs(ids []int) ([]user_types.User, error)
}

type Controller struct {
	service IService
}

func NewController(service IService) *Controller {
	return &Controller{
		service: service,
	}
}