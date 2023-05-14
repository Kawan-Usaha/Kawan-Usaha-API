package main

import (
	"fmt"
	"kawan-usaha-api/server"
	"log"
)

func main() {
	fmt.Println("Hello world")

	r := server.SetupRouter()

	//Server init

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
		return
	}
}
