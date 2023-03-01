package server

import (
	"http2/app/controller"
	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
) 

func Server() *gin.Engine{
	godotenv.Load()
	router := gin.Default()
	router.GET("/getUser", controller.GetUser)
	router.POST("/signIn", controller.SignIn)
	router.GET("/parse", controller.Parse)
	router.POST("/token", controller.ParseBearer)

	return router
}