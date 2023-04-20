package purchases_controller

import (
	"http2/app/types/purchases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *PurchController) Purchases(c *gin.Context) {
	var purch purchases_type.Purchases
	if err := c.BindJSON(&purch); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err := controller.service.Purchases(purch.IdBuyer, purch.IdGoods)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"purchase": "OK"})
	return
}

func (controller *PurchController) GetPurchases(c *gin.Context) {
	data, err := controller.service.GetPurchases()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}