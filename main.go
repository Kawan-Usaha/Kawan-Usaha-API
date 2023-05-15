package main

import (
	"fmt"
	"kawan-usaha-api/server"
	"kawan-usaha-api/server/lib"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello world")

	// Load .env file
	lib.EnvLoader()

	r := server.SetupRouter()

	var result string
	var err error

	if result, err = lib.GenerateToken(time.Duration(30), "123123123 ini username", "t"); err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(result)

	var results interface{}
	if results, err = lib.ValidateToken(result, "t"); err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(results)

	//Server init
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
		return
	}
}
