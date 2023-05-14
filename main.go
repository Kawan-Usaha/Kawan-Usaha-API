package main

import (
	"fmt"
	"kawan-usaha-api/server"
	"kawan-usaha-api/server/lib"
	"log"
)

func main() {
	fmt.Println("Hello world")

	// Load .env file
	lib.EnvLoader()

	r := server.SetupRouter()

	//Server init

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
		return
	}
}
