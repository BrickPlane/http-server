package service

import (
	"http2/app/types"

	"github.com/jmoiron/sqlx"
)

type IStorProd interface {
	SaveProduct(val types.SaveProductsRequest) (*types.ProductsResponce, error)
	GetProduct() ([]types.ProductsResponce, error)
	GetProductByID(id uint64) (*types.ProductsResponce, error)
	UpdateProd(val types.UpdateProductsRequest) (*types.ProductsResponce, error)
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
