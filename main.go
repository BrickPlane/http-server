package main

import (
	"fmt"
	"http2/app/controller/product"
	purchases_controller "http2/app/controller/purchases"
	"http2/app/controller/user"
	"http2/app/controller/wallet"
	"http2/app/server"
	"http2/app/service/product"
	purchases_service "http2/app/service/purchase"
	"http2/app/service/user"
	"http2/app/service/wallet"
	"http2/app/storage/product"
	purch_storage "http2/app/storage/purchases"
	"http2/app/storage/user"
	"http2/app/storage/wallet"

	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("secret.env")

	userDB, err := user_storage.NewStorage()
	if err != nil {
		fmt.Println(" Get error while creating DB connect ", err)
		os.Exit(1)
	}

	productDB, err := product_storage.NewProdStorage()
	if err != nil {
		fmt.Println(" Get error while creating DB connect ", err)
		os.Exit(1)
	}

	walletDB, err := wallet.NewWalletStorage()
	if err != nil {
		fmt.Println(" Get error while creating DB connect ", err)
		os.Exit(1)
	}

	purchasesDB, err := purch_storage.NewStorage()
	if err != nil {
		fmt.Println(" Get error while creating DB connect ", err)
		os.Exit(1)
	}

	
	productService := product_service.NewProdService(productDB, productDB.DB)
	userService := user_service.NewService(userDB, userDB.DB)
	walletService := wallet_service.NewService(walletDB, walletDB.DB)
	purchasesService := purchases_service.NewPurchService(purchasesDB, userDB, productDB, purchasesDB.DB)
	userController := user_controller.NewController(userService)
	productController := product_controller.NewProdController(productService)
	walletController := wallet_controller.NewController(walletService)
	purchasesController := purchases_controller.NewPurchController(purchasesService)
	router := server.Server(userController, productController, walletController, purchasesController)
	
	router.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}