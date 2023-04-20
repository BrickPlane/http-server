package product_service

import (
	"errors"
	"http2/app/types/productDB"
)

func (srv *ProdService) AddProduct(catalog product_types.SaveProductsRequest) (*product_types.ProductsResponce, error) {
	data, err := srv.storageProd.SaveProduct(catalog)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *ProdService) GetProduct() ([]product_types.ProductsResponce, error) {
	data, err := srv.storageProd.GetProduct()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *ProdService) GetProductByID(id uint64) (*product_types.ProductsResponce, error) {
	data, err := srv.storageProd.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *ProdService) UpdateProduct(upd product_types.UpdateProductsRequestDTO) (*product_types.ProductsResponce, error) {
	varProduct := make(map[string]interface{})

	if upd.Brand != nil && len(*upd.Brand) !=0 {
		varProduct["brand"] = *upd.Brand
	}

	if upd.Name != nil && len(*upd.Name) !=0 {
		varProduct["name"] = *upd.Name
	}

	if upd.Price != nil && *upd.Price <=0 {
		varProduct["price"] = *upd.Price
	}

	if len(varProduct) == 0 {
		return nil, errors.New("Nothing to changed")
	}

	data, err := srv.storageProd.UpdateProd(upd.ID, varProduct)
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
