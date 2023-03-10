package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type IController interface {
	SignIn(c *gin.Context)
	ParseBearer(c *gin.Context)
	GetUser(c *gin.Context)
}

func Server(controller IController) *gin.Engine {
	godotenv.Load()
	router := gin.Default()
	router.POST("/signIn", controller.SignIn)
	router.POST("/token", controller.ParseBearer)
	router.GET("/getUser", controller.GetUser)

	return router
}
