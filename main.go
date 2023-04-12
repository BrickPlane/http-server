package main

import (
	"fmt"
	"http2/app/controller"
	"http2/app/server"
	"http2/app/service"
	"http2/app/storage"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("secret.env")

	startDB, err := storage.NewStorage()
	if err != nil {
		fmt.Println(" Get error while creating DB connect ", err)
		os.Exit(1)
	}

	prodStorage, err := storage.NewProdStorage()
	if err != nil {
		fmt.Println(" Get error while creating DB connect ", err)
		os.Exit(1)
	}

	
	prodService := service.NewProdService(prodStorage, prodStorage.DB)
	userService := service.NewService(startDB, startDB.DB)
	userController := controller.NewController(userService)
	prodController := controller.NewProdController(prodService)
	router := server.Server(userController, prodController)
	
	router.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}