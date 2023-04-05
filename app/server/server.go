package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type IController interface {
	HandlerFunc() gin.HandlerFunc
	Signin(c *gin.Context)
	Login(c *gin.Context)
	ParseBearer(c *gin.Context)
	GetAllUser(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetUserByLogin(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetUserByIDs(c *gin.Context)
}

func Server(controller IController) *gin.Engine {
	godotenv.Load()
	router := gin.Default()
	router.POST("/registration", controller.Signin)
	router.POST("/getToken", controller.ParseBearer)
	router.GET("/getAllUsers", controller.GetAllUser)
	router.GET("/GetUserByIDs", controller.GetUserByIDs)
	router.GET("/login", controller.Login)


	// apiRouters := router.Group("/api", controller.HandlerFunc())
	router.Use(controller.HandlerFunc())
	router.GET("/getUserByID", controller.GetUserByID)
	router.GET("/getUser", controller.GetUserByLogin)
	router.PATCH("/update", controller.Update)
	router.DELETE("/delete", controller.Delete)

	return router
}
