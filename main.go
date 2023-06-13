package main

import (
	"kawan-usaha-api/server"
	"kawan-usaha-api/server/lib"
	"log"
)

func main() {
	log.Println("Starting server...")

	// Load .env file
	lib.EnvLoader()

	r := server.SetupRouter()

	//Server init
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err.Error())
		return
	}
}
