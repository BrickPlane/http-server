package service

import (
	"http2/app/types"
)

func (srv *ProdService) AddProduct(catalog types.SaveProductsRequest) (*types.ProductsResponce, error) {
	data, err := srv.storageProd.SaveProduct(catalog)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *ProdService) GetProduct() ([]types.ProductsResponce, error) {
	data, err := srv.storageProd.GetProduct()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *ProdService) GetProductByID(id uint64) (*types.ProductsResponce, error) {
	data, err := srv.storageProd.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *ProdService) UpdateProduct(catalog types.UpdateProductsRequest) (*types.ProductsResponce, error) {
	data, err := srv.storageProd.UpdateProd(catalog)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *ProdService) DeleteProduct(id uint64) error {
	err := srv.storageProd.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
