package lib

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvLoader() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err.Error())
	}
}

func EnvLoaderTest() {
	// Load .env file for test
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err.Error())
	}
}
