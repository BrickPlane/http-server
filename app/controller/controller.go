package controller

import (
	"net/http"

	"http2/app/storage"

	"github.com/gin-gonic/gin"
)

type IService interface {
	SignToken(c *gin.Context, creds storage.Credential) (string, error)
}

type Controller struct {
	service IService
}

func NewController(service IService) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller)SignIn(c *gin.Context) {
	var creds storage.Credential
	if err := c.BindJSON(&creds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Wrong input data")
		return
	}

	// token, err := service.SignToken(c, creds)
	token, err := controller.service.SignToken(c, creds)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, token)
}
