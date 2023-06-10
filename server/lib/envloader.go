package lib

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoader() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Deployment mode is: %s", os.Getenv("DEPLOYMENT_MODE"))
}

func EnvLoaderTest() {
	// Load .env file for test
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err.Error())
	}
}
