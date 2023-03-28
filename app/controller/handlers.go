package controller

import (
	"http2/app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds types.Credential

		if err := c.BindJSON(&creds); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "[hendlers] Wrong input data")
			return
		}
		if err := controller.service.ParseWithBearer(c, creds); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		c.Set("key", creds)
		c.Next()
		return
	}
}

func (controller *Controller) Signin(c *gin.Context) {
	var creds types.Credential
	if err := c.BindJSON(&creds); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Wrong input data [signIn]")
		return
	}

	data, err := controller.service.SigninUser(creds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	token, _ := controller.service.GenToken(c, creds)
	var userInfo []any
	userInfo = append(userInfo, data, token)
	c.IndentedJSON(http.StatusOK, userInfo)
	return
}

func (controller *Controller) ParseBearer(c *gin.Context) {
	// err := controller.service.ParseWithBearer(c)
	// err = nil
	// if err = nil {
	// 	// c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	// 	return
	// }
	c.IndentedJSON(http.StatusOK, gin.H{"msg": "func not used now"})
	return
}

func (controller *Controller) GetAllUser(c *gin.Context) {
	allUser, err := controller.service.GetAllUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, allUser)
	return
}

func (controller *Controller) GetUser(c *gin.Context) {
	getKey := c.MustGet("key")
	get, ok := getKey.(types.Credential)
	if !ok {
		return
	}

	idUser, err := controller.service.GetUser(get)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, idUser)
	return
}

type GetUserByIDRequest struct {
	IDs []int `json:"ids"`
}

func (controller *Controller) GetUserByIDs(c *gin.Context) {
	var req GetUserByIDRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Wrong input data")
		return
	}

	users, err := controller.service.GetUserByIDs(req.IDs)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, users)
	return
}

func (Controller *Controller) Update(c *gin.Context) {
	updKey := c.MustGet("key")
	upd, ok := updKey.(types.Credential)
	if !ok {
		return
	}

	data, err := Controller.service.UpdateUser(upd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, data)
	return
}

func (Controller *Controller) Delete(c *gin.Context) {
	dltKey := c.MustGet("key")
	dlt, ok := dltKey.(types.Credential)
	if !ok {
		return
	}
	err := Controller.service.DeleteUser(dlt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "User deleted")
	return
}
