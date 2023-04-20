package product_storage

import (
	"database/sql"
	"fmt"

	"http2/app/storage"
	"http2/app/types/errors"
	"http2/app/types/productDB"

	"github.com/jmoiron/sqlx"
)

type ProdStorage struct {
	DB *sqlx.DB
}

func NewProdStorage() (*ProdStorage, error) {
	con, err := storage.ConnectDB()
	if err != nil {
		return nil, err
	}

	return &ProdStorage{DB: con}, nil
}

func (stor *ProdStorage) SaveProduct(val product_types.SaveProductsRequest) (*product_types.ProductsResponce, error) {
	query, err := stor.DB.Query(`
	INSERT INTO "product" (name, brand, price)
	VALUES ($1, $2, $3)
	RETURNING *`,
		val.Name, val.Brand, val.Price,
	)
	if err != nil {
		return nil, err
	}

	data := &product_types.ProductsResponce{}
	for query.Next() {
		if err := query.Scan(
			&data.ID,
			&data.Name,
			&data.Brand,
			&data.Price,
		); err != nil {
			return nil, errors.CredError
		}
	}

	return data, nil
}

func (stor *ProdStorage) GetProduct() ([]product_types.ProductsResponce, error) {
	query, err := stor.DB.Query(`SELECT * FROM "product"`)
	if err != nil {
		return nil, err
	}

	slice := make([]product_types.ProductsResponce, 0)
	for query.Next() {
		var data product_types.ProductsResponce
		if err := query.Scan(
			&data.ID,
			&data.Name,
			&data.Brand,
			&data.Price,
		); err != nil {
			return nil, errors.CredError
		}
		slice = append(slice, data)
	}

	return slice, nil
}

func (stor *ProdStorage) GetProductByID(id uint64) (*product_types.ProductsResponce, error) {
	data := &product_types.ProductsResponce{}

	err := stor.DB.Get(data, `SELECT * FROM "product" WHERE id=$1`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return data, nil
}

func (stor *ProdStorage) DeleteProduct(id uint64) error {
	_, err := stor.DB.Query(`DELETE FROM "product" WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (stor *ProdStorage) UpdateProd(id int, val map[string]interface{}) (*product_types.ProductsResponce, error) {
	queryForUpdate := storage.HelperForUpdate(val)

	query := fmt.Sprintf(`UPDATE "product" SET %[1]s WHERE id=%[2]d RETURNING *`,
	queryForUpdate, id)
	data, err := stor.DB.NamedQuery(query, val)
	if err != nil {
		return nil, err
	}

	scan := &product_types.ProductsResponce{}
	for data.Next() {
		if err := data.Scan(
			&scan.ID,
			&scan.Name,
			&scan.Brand,
			&scan.Price,
		); err != nil {
			return nil, errors.CredError
		}
	}
	return scan, nil
}