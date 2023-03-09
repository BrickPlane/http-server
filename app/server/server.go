package server

import (

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
) 

type IController interface{
	SignIn(c *gin.Context)
	ParseBearer(c *gin.Context)
}

func Server(controller IController) *gin.Engine{
	godotenv.Load()
	router := gin.Default()
	router.POST("/signIn", controller.SignIn)
	// router.Use(service.HandlerFunc())
	router.POST("/token", controller.ParseBearer)

	return router
}


func test() *gin.Engine {
	j := gin.Default()
	// j.POST("/aaa", controller.ParseBearer)
	return j
}