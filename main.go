package main

import (
	"os"

	"http2/app/server"

	"github.com/joho/godotenv"
) 

func main() {
	godotenv.Load("secret.env")
	router := server.Server()
	router.Run(os.Getenv("HOST")+":"+os.Getenv("PORT"))
}

