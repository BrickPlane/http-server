package types

type ProductsResponce struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Brand string `json:"brand" db:"brand"`
	Price string `json:"price" db:"price"`
}

type ProductsRequest struct {
	Name  string `json:"name" db:"name"`
	Brand string `json:"brand" db:"brand"`
	Price string `json:"price" db:"price"`
}

type SaveProductsRequest struct {
	ProductsRequest
}

type UpdateProductsRequest struct {
	ID    int    `json:"id" db:"id"`
	Name  *string `json:"name" db:"name"`
	Brand *string `json:"brand" db:"brand"`
	Price *string `json:"price" db:"price"`
}
