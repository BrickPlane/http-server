package wallet_service

import (
	"http2/app/types/walletDB"
)

func (service *Service) ReplenishmentWallet(val wallet_types.Replenishment) error {
	err := service.storage.ReplenishmentWallet(val)
	if err != nil {
		return err
	}
	return nil
}