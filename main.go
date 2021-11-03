package main

import (
	"github.com/Ad3bay0c/ContactSaver/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("There was an error loading.env file")
	}

	newServer := &server.Server{}

	newServer.Start()
}
