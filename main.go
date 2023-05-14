package main

import (
	"fmt"
	"kawan-usaha-api/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	r := server.SetupRouter()

	//Server init

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
		return
	}
}
