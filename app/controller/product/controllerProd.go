package product_controller

import (
	"http2/app/types/errors"
	"http2/app/types/productDB"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *ProdController) AddProduct(c *gin.Context) {
	var catalog product_types.SaveProductsRequest
	if err := c.BindJSON(&catalog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.InputData)
		return
	}

	data, err := controller.service.AddProduct(catalog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
	return
}

func (controller *ProdController) GetProduct(c *gin.Context) {
	data, err := controller.service.GetProduct()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
	return
}

func (controller *ProdController) GetProductByID(c *gin.Context) {
	var catalog product_types.ProductID
	if err := c.BindJSON(&catalog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	data, err := controller.service.GetProductByID(catalog.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
	return
}

func (controller *ProdController) DeleteProduct(c *gin.Context) {
	var catalog product_types.ProductID
	if err := c.BindJSON(&catalog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err := controller.service.DeleteProduct(catalog.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, "product exterminated")
	return
}

func (controller *ProdController) UpdateProduct(c *gin.Context) {
	var catalog product_types.UpdateProductsRequestDTO
	if err := c.BindJSON(&catalog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	data, err := controller.service.UpdateProduct(catalog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
	return 
}
