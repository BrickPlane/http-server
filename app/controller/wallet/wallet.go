package wallet_controller

import (
	"fmt"
	"http2/app/types/errors"
	"http2/app/types/walletDB"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) ReplenishmentWallet(c *gin.Context) {
	var money wallet_types.Replenishment
	if err := c.BindJSON(&money); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.InputData)
		return
	}

	if err := controller.service.ReplenishmentWallet(money); err != nil {
		fmt.Println("err wallet",err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"the account was successfully replenished by the amount of ":money.Fill})
	return
}
