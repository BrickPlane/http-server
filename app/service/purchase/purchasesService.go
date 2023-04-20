package purchases_service

import (
	purchases_type "http2/app/types/purchases"
)

func (service *PurchService) Purchases(usr uint64, prdct uint64) error{
	user, product, err := service.PurchHelper(usr, prdct)
	if err != nil {
		return err
	}
	
	trx, err := service.storagePurch.Transaction()
	if err != nil {
		return err
	}

	err = service.storagePurch.ChangeWallet(user.Wallet-product.Price, user.ID, trx)
	if err!= nil {
		return err
	}

	err = service.storagePurch.Receipt(user, product, trx)
	if err != nil {
		return err
	}
	return nil
}

func (service *PurchService) GetPurchases() ([]purchases_type.Purchases, error){
	data, err := service.storagePurch.GetPurchased()
	if err != nil {
		return nil, err
	}
	return data, nil
}
