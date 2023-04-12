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
	UpdateUser(upd types.UpdateUserRequestDTO) (*types.UpdateUserResponseDTO, error)
	DeleteUser(dlt uint64) error
	GetUserByID(id uint64) (*types.User, error)
	GetUserByIDs(ids []int) ([]types.User, error)

	// AddProduct(catalog types.ProductsRequest) (*types.ProductsResponce, error)
	// GetProduct() ([]types.ProductsResponce, error)
	// GetProductByID(id uint64) (*types.ProductsResponce, error)
	// UpdateProduct(catalog types.ProductsRequest) (*types.ProductsResponce, error)
	// DeleteProduct(id uint64) error
}

type Controller struct {
	service IService
}

func NewController(service IService) *Controller {
	return &Controller{
		service: service,
	}
}