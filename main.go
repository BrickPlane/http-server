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


// TODO connect to database ✓
// TODO create table if didn`t exist ✓
// TODO rout signIn -> write in DB ✓
// TODO make getAllUser route - interface + DB ✓ 

// TODO User can create account in the system (Save in DB email and password)
// TODO User can update account (email or password. If email already exist return error message)
// TODO User can get info about him (Send ID or email and server has return info about user)
// TODO User can delete his account (Can delete only his own account)
// TODO Login to the system. If user was created in the system and his credential is correct, send JWT token.
// TODO Fresh and correct JWT token give user do points 2 - 4 
