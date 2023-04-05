package controller

import (
	"fmt"
	"http2/app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResp struct {
	ErrorReason *string `json:"err"`
	Data        *any    `json:"data"`
}

func (controller *Controller) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req interface{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		c.Set("key", req)

		token, err := controller.service.ParseWithBearer(c)
		if err == nil {
			fmt.Println(err)
			// c.Next()
			return
		}

		err = controller.service.TokenVerification(token)
		if err != nil {
			fmt.Println("[handler] check", err)
			fmt.Println("[handler] token", token)
			errResp := new(CommonResp)
			reason := err.Error()
			errResp.ErrorReason = &reason
			fmt.Println("[handler] token", errResp)

			c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
			c.Next()
			return
		}

		formatData, err := convertToType[any, types.Credential](req)
		if err != nil {
			fmt.Println("error2", err)

			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		if _, err := controller.service.Login(*formatData); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			fmt.Println("error3", err)

			c.Next()
			return
		}

	}
}

func (controller *Controller) Signin(c *gin.Context) {
	var info types.User
	if err := c.BindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Wrong input data [signIn]")
		return
	}

	data, err := controller.service.SigninUser(info)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	token, err := controller.service.GenToken(c, *data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "query error")
		return
	}
	var userInfo []any
	userInfo = append(userInfo, data, token)
	c.IndentedJSON(http.StatusOK, userInfo)
	return
}

func (controller *Controller) Login(c *gin.Context) {
	var creds types.Credential
	if err := c.BindJSON(&creds); err != nil {
		_, err := controller.service.ParseWithBearer(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}
	data, err := controller.service.Login(creds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
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

func (controller *Controller) GetUserByID(c *gin.Context) {
	data, exist := c.Get("key")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, "error during data transfer")
		return
	}

	formatData, err := convertToType[any, types.UserID](data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	idUser, err := controller.service.GetUserByID(formatData.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, idUser)
	return
}

func (Controller *Controller) GetUserByLogin(c *gin.Context) {
	data, exist := c.Get("key")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, "error during data transfer")
		return
	}

	formatData, err := convertToType[any, types.Credential](data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	idUser, err := Controller.service.GetUserByLogin(formatData.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, idUser)
	return
}

func (Controller *Controller) Update(c *gin.Context) {
	updKey, exist := c.Get("key")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, "error during data transfer")
		return
	}

	formatData, err := convertToType[any, types.User](updKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	data, err := Controller.service.UpdateUser(*formatData) // ??
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, data)
	return
}

func (Controller *Controller) Delete(c *gin.Context) {
	data, exist := c.Get("key")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, "error during data transfer")
		return
	}

	formatData, err := convertToType[any, types.UserID](data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	er := Controller.service.DeleteUser(formatData.ID)
	if er != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, er)
		return
	}

	c.IndentedJSON(http.StatusOK, "User deleted")
	return
}
