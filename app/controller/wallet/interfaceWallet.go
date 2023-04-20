package wallet_controller

import (
	"http2/app/types/walletDB"
)

type IService interface {
	ReplenishmentWallet(val wallet_types.Replenishment) error
}

type Controller struct {
	service IService
}

func NewController(service IService) *Controller {
	return &Controller{
		service: service,
	}
}