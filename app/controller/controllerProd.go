package controller

import (
	"fmt"
	"http2/app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *ProdController) AddProduct(c *gin.Context) {
	var catalog types.SaveProductsRequest
	if err := c.BindJSON(&catalog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	// _, err := controller.service.ParseWithBearer(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, err)
	// 	return
	// }

	data, err := controller.service.AddProduct(catalog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (controller *ProdController) GetProduct(c *gin.Context) {
	data, err := controller.service.GetProduct()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (controller *ProdController) GetProductByID(c *gin.Context) {
	var catalog types.UserID
	if err := c.BindJSON(&catalog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	// _, err := controller.service.ParseWithBearer(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, err)
	// 	return
	// }

	data, err := controller.service.GetProductByID(catalog.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (controller *ProdController) DeleteProduct(c *gin.Context) {
	var catalog types.UserID
	if err := c.BindJSON(&catalog); err != nil {
		fmt.Println("error 1", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	// _, err := controller.service.ParseWithBearer(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, err)
	// 	return
	// }

	err := controller.service.DeleteProduct(catalog.ID)
	if err != nil {
		fmt.Println("error 2", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, "product exterminated")
	return
}

func (controller *ProdController) UpdateProduct(c *gin.Context) {
	var catalog types.UpdateProductsRequest
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
