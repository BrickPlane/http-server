package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IService interface {
	SignToken(c *gin.Context) string
}

type Controller struct {
	service IService
}

func NewController(service2 IService) *Controller {
	
	return &Controller{
		service: service2,
	}
}

func (controller *Controller)SignIn(c *gin.Context) {
	token := controller.service.SignToken(c)
	c.IndentedJSON(http.StatusOK, token)
}
