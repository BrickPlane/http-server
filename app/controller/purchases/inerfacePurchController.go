package purchases_controller

import  (
	"http2/app/types/userDB"
	"http2/app/types/productDB"
	"http2/app/types/purchases"
)

type IServicePurch interface {
	Purchases(usr uint64, prdct uint64) error
	GetPurchases() ([]purchases_type.Purchases, error)
	PurchHelper(usr uint64, prdct uint64) (*user_types.User, *product_types.ProductsResponce, error)
}

type PurchController struct {
	service IServicePurch
}

func NewPurchController(service IServicePurch) *PurchController {
	return &PurchController{
		service: service,
	}
}