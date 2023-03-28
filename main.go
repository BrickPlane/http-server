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
	
	service := service.NewService(startDB, startDB.DB)
	controller := controller.NewController(service)
	router := server.Server(controller)
	
	router.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}



// TODO User can update account (email or password. If email already exist return error message)
// TODO User can get info about him (Send ID or email and server has return info about user)
// TODO User can delete his account (Can delete only his own account)
// TODO Fresh and correct JWT token give user do points 1 - 3 
