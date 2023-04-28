package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

type IWalletController interface {
	ReplenishmentWallet(c *gin.Context)
}

type IPurchasesController interface {
	Purchases(c *gin.Context)
	GetPurchases(c *gin.Context)
}

func Server(
	controller IUserController, 
	prodController IProdController, 
	walletController IWalletController,
	purchasesController IPurchasesController,
	) *gin.Engine {
	godotenv.Load()
	router := gin.Default()

	userRouters := router.Group("/user")
	productRouters := router.Group("/product")
	purchasesRouters := router.Group("/purchases")

	router.POST("/fill", walletController.ReplenishmentWallet)
	
	userRouters.POST("/registration", controller.Signin)
	userRouters.POST("/getToken", controller.ParseBearer)
	userRouters.GET("/getAllUsers", controller.GetAllUser)
	userRouters.GET("/GetUserByIDs", controller.GetUserByIDs)
	userRouters.GET("/login", controller.Login)

	productRouters.POST("/add", prodController.AddProduct)
	productRouters.GET("/get", prodController.GetProduct)
	productRouters.GET("/getID", prodController.GetProductByID)
	productRouters.DELETE("/deleteProductByID", prodController.DeleteProduct)
	productRouters.PATCH("/updateProduct", prodController.UpdateProduct)

	purchasesRouters.POST("/buy", purchasesController.Purchases)
	purchasesRouters.GET("/all", purchasesController.GetPurchases)

	userRouters.Use(controller.HandlerFunc())
	userRouters.GET("/getUserByID", controller.GetUserByID)
	userRouters.GET("/getUser", controller.GetUserByLogin)
	userRouters.PATCH("/update", controller.Update)
	userRouters.DELETE("/delete", controller.Delete)

	return router
}