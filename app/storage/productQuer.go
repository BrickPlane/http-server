package storage

import (
	"database/sql"
	"fmt"

	"http2/app/types"

	"github.com/jmoiron/sqlx"
)

type ProdStorage struct {
	DB *sqlx.DB
}

func NewProdStorage() (*ProdStorage, error) {
	con, err := connectDB()
	if err != nil {
		return nil, err
	}

	return &ProdStorage{DB: con}, nil
}

func (stor *ProdStorage) SaveProduct(val types.SaveProductsRequest) (*types.ProductsResponce, error) {
	query, err := stor.DB.Query(`
	INSERT INTO "product" (name, brand, price)
	VALUES ($1, $2, $3)
	RETURNING *`,
		val.Name, val.Brand, val.Price,
	)
	if err != nil {
		return nil, err
	}

	data := &types.ProductsResponce{}
	for query.Next() {
		if err := query.Scan(
			&data.ID,
			&data.Name,
			&data.Brand,
			&data.Price,
		); err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (stor *ProdStorage) GetProduct() ([]types.ProductsResponce, error) {
	query, err := stor.DB.Query(`SELECT * FROM "product"`)
	if err != nil {
		return nil, err
	}

	slice := make([]types.ProductsResponce, 0)
	for query.Next() {
		var data types.ProductsResponce
		if err := query.Scan(
			&data.ID,
			&data.Name,
			&data.Brand,
			&data.Price,
		); err != nil {
			return nil, err
		}
		slice = append(slice, data)
	}

	return slice, nil
}

func (stor *ProdStorage) GetProductByID(id uint64) (*types.ProductsResponce, error) {
	data := &types.ProductsResponce{}
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

func (stor *ProdStorage) UpdateProd(val types.UpdateProductsRequest) (*types.ProductsResponce, error) {
	data, err := stor.DB.Query(`UPDATE "product" SET name=$1, brand=$2, price=$3 WHERE id=$4 RETURNING *`, val.Name, val.Brand, val.Price, val.ID)
	if err != nil {
		return nil, err
	}

	scan := &types.ProductsResponce{}
	for data.Next() {
		if err := data.Scan(
			&scan.ID,
			&scan.Name,
			&scan.Brand,
			&scan.Price,
		); err != nil {
			return nil, err
		}
	}
	fmt.Println("scan", scan)
	fmt.Println("data", data)
	fmt.Println("val", val)

	return scan, nil
}