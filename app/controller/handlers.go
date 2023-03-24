package controller

import (
	"http2/app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) Signin(c *gin.Context) {
	var creds types.Credential
	if err := c.BindJSON(&creds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Wrong input data")
		return
	}

	data, err := controller.service.SigninUser(creds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return

	}
	c.IndentedJSON(http.StatusOK, data)
	return
	// 	var newUser storage.Credential
	// 	if err := c.BindJSON(&newUser); err != nil {
	// 		c.IndentedJSON(http.StatusBadRequest, "wrong input data")
	// 		return
	// 	}
	// 	q := `INSERT INTO credential (login, password) VALUES (:login, :password) ON CONFLICT DO NOTHING`
	// 	result, err := controller.DB.QueryRow(q, newUser.Login, newUser.Password)
	// 	if err != nil {
	// 		c.IndentedJSON(http.StatusInternalServerError, "can`t record new data")
	// 	}
	// 	n, err := result.Columns()
	// 	if err != nil {
	// 		c.IndentedJSON(http.StatusInternalServerError, "error check rows DB")
	// 	}
	// 	if n == nil {
	// 		c.IndentedJSON(http.StatusInternalServerError, "can`t record data")
	// 	}
	// 	c.IndentedJSON(http.StatusOK, "successfuly create user")
	// }

	// func (controller *Controller) Login(c *gin.Context) {
	// 	var creds storage.Credential
	// 	if err := c.BindJSON(&creds); err != nil {
	// 		c.IndentedJSON(http.StatusBadRequest, "Wrong input data")
	// 		return
	// 	}
	// 	token, err := controller.service.LoginToken(c, creds)
	// 	if err != nil {
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// 	for _, a := range storage.Users {
	// 		if a.Login == creds.Login {
	// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "user already exist"})
	// 			return
	// 		}
	// 	}
	// storage.Users = append(storage.Users, creds)
	// c.IndentedJSON(http.StatusOK, gin.H{"token:": token})
}

func (controller *Controller) ParseBearer(c *gin.Context) {
	err := controller.service.ParseWithBearer(c)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"msg": "Token valid"})
}

func (controller *Controller) GetUser(c *gin.Context) {
	allUser, err := controller.service.GetUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, allUser)
	return
}

func (controller *Controller) GetUserByID(c *gin.Context) {
	var get types.Credential
	if err := c.BindJSON(&get); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Wrong input data")
		return
	}

	idUser, err := controller.service.GetUserByID(get)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, idUser)
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
}

func (Controller *Controller) Update(c *gin.Context) {
	var upd types.Credential
	if err := c.BindJSON(&upd); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Wrong input data")
		return
	}

	data, err := Controller.service.UpdateUser(upd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
	return
}

func (Controller *Controller) Delete(c *gin.Context) {
	var dlt types.Credential
	if err := c.BindJSON(&dlt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Wrong input data")
		return
	}

	err := Controller.service.DeleteUser(dlt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, "User deleted")
	return
}
