package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type IController interface {
	Signin(c *gin.Context)
	ParseBearer(c *gin.Context)
	GetUser(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetUserByIDs(c *gin.Context)
}

func Server(controller IController) *gin.Engine {
	godotenv.Load()
	router := gin.Default()
	router.POST("/registration", controller.Signin)
	router.POST("/getToken", controller.ParseBearer)
	router.GET("/getUsers", controller.GetUser)
	router.GET("/getUserByID", controller.GetUserByID)
	router.GET("/GetUserByIDs", controller.GetUserByIDs)
	router.PATCH("/update", controller.Update)
	router.DELETE("/delete", controller.Delete)
	return router
}
