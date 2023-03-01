package controller

import (
	"net/http"

	"http2/app/service"
	"http2/app/storage"

	"github.com/gin-gonic/gin"
) 

func GetUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.Users)
}

func SignIn(c *gin.Context) {
	var creds storage.Credential
	if err := c.BindJSON(&creds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Wrong input data")
		return
	}

	token, err := service.SignToken(c, creds);
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, token)		
}

func Parse(c *gin.Context) {
	var token string 
	if err := c.BindJSON(&token); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Wrong input token data")
		return
	}

	jwtToken, err := service.Parse(c, token);
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}
	c.IndentedJSON(http.StatusOK, jwtToken)
}

func ParseBearer(c *gin.Context) {
	if err := service.WithBearer(c); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}
}