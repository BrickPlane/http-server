package product_controller

import (
	"http2/app/types/productDB"

)


type IServiceProd interface {
	AddProduct(catalog product_types.SaveProductsRequest) (*product_types.ProductsResponce, error)
	GetProduct() ([]product_types.ProductsResponce, error)
	GetProductByID(id uint64) (*product_types.ProductsResponce, error)
	UpdateProduct(catalog product_types.UpdateProductsRequestDTO) (*product_types.ProductsResponce, error)
	DeleteProduct(id uint64) error
}

type ProdController struct {
	service IServiceProd
}

func NewProdController(service IServiceProd) *ProdController {
	return &ProdController{
		service: service,
	}
}
