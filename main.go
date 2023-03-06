package main

import (
	"http2/app/controller"
	"http2/app/server"
	"http2/app/service"
	"os"

	"github.com/joho/godotenv"
) 

func main() {
	godotenv.Load("secret.env")
	service := service.NewService()
	controller := controller.NewController(service)
	router := server.Server(controller)
	router.Run(os.Getenv("HOST")+":"+os.Getenv("PORT"))	
}