package product_service

import (
	"http2/app/types/productDB"

	"github.com/jmoiron/sqlx"
)

type IStorProd interface {
	SaveProduct(val product_types.SaveProductsRequest) (*product_types.ProductsResponce, error)
	GetProduct() ([]product_types.ProductsResponce, error)
	GetProductByID(id uint64) (*product_types.ProductsResponce, error)
	UpdateProd(id int, val map[string]interface{}) (*product_types.ProductsResponce, error)
	DeleteProduct(id uint64) error
}

type ProdService struct {
	DB          *sqlx.DB
	storageProd IStorProd
}

func NewProdService(storage IStorProd, conn *sqlx.DB) *ProdService {
	return &ProdService{
		storageProd: storage,
		DB:          conn,
	}
}
