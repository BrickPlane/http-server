package purchases_service

import (
	"errors"
	"http2/app/types/userDB"
	"http2/app/types/productDB"
)

func (service *PurchService) PurchHelper(usr uint64, prdct uint64) (*user_types.User, *product_types.ProductsResponce, error){
	user, _ := service.storageUser.GetUserByID(usr)
	if user.Wallet <= 0 {
		return nil, nil, errors.New("fill the bill") 
	}

	product, err := service.storageProduct.GetProductByID(prdct)
	if err != nil {
 		return nil, nil, errors.New("the product is out of stock")
	}

	if user.Wallet < product.Price {
		return nil, nil, errors.New("not enough money")
	}
	return user, product, nil
}