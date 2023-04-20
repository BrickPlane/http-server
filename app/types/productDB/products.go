package product_types

type ProductsResponce struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Brand string `json:"brand" db:"brand"`
	Price float64 `json:"price" db:"price"`
}

type ProductsRequest struct {
	Name  string `json:"name" db:"name"`
	Brand string `json:"brand" db:"brand"`
	Price int `json:"price" db:"price"`
}

type SaveProductsRequest struct {
	ProductsRequest
}

type UpdateProductsRequestDTO struct {
	ID    int     `json:"id" db:"id"`
	Name  *string `json:"name" db:"name"`
	Brand *string `json:"brand" db:"brand"`
	Price *float64 `json:"price" db:"price"`
}

type CommonRequest struct {
	ProductID
}

type ProductID struct {
	ID uint64 `json:"id" db:"id"`
}
