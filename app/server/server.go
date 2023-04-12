package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// type IController interface {
// 	IUserController
// 	IProdController
// }
type IUserController interface {
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

type IProdController interface {
	AddProduct(c *gin.Context)
	GetProduct(c *gin.Context)
	GetProductByID(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

func Server(controller IUserController, prodController IProdController) *gin.Engine {
	godotenv.Load()
	router := gin.Default()
	router.POST("/registration", controller.Signin)
	router.POST("/getToken", controller.ParseBearer)
	router.GET("/getAllUsers", controller.GetAllUser)
	router.GET("/GetUserByIDs", controller.GetUserByIDs)
	router.GET("/login", controller.Login)

	router.POST("/add", prodController.AddProduct)
	router.GET("/get", prodController.GetProduct)
	router.GET("/getID", prodController.GetProductByID)
	router.DELETE("/deleteProductByID", prodController.DeleteProduct)
	router.PATCH("/updateProduct", prodController.UpdateProduct)

	// apiRouters := router.Group("/api", controller.HandlerFunc())
	router.Use(controller.HandlerFunc())
	router.GET("/getUserByID", controller.GetUserByID)
	router.GET("/getUser", controller.GetUserByLogin)
	router.PATCH("/update", controller.Update)
	router.DELETE("/delete", controller.Delete)

	return router
}
