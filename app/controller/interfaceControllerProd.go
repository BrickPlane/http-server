package controller

import (
	"http2/app/types"

	// "github.com/gin-gonic/gin"
)


type IServiceProd interface {
	AddProduct(catalog types.SaveProductsRequest) (*types.ProductsResponce, error)
	GetProduct() ([]types.ProductsResponce, error)
	GetProductByID(id uint64) (*types.ProductsResponce, error)
	UpdateProduct(catalog types.UpdateProductsRequest) (*types.ProductsResponce, error)
	DeleteProduct(id uint64) error

	// ParseWithBearer(c *gin.Context) (*types.JWTUploadData, error)
}

type ProdController struct {
	service IServiceProd
}

func NewProdController(service IServiceProd) *ProdController {
	return &ProdController{
		service: service,
	}
}
