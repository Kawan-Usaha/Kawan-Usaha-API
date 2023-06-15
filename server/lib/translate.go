package lib

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/translate"
	"google.golang.org/api/option"
)

var once sync.Once

func createTranslationClient() (*translate.Client, error) {
	ctx := context.Background()
	// Retrieve the path to the service account key file from the environment variable
	serviceAccountKeyPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	// Create the client with the specified service account key file
	client, err := translate.NewClient(ctx, option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	return client, nil
}

func GetTranslationClient() (*translate.Client, error) {
	var client *translate.Client
	var err error
	once.Do(func() {
		client, err = createTranslationClient()
		if err != nil {
			log.Panicf("Failed to create client: %v", err)
			return
		}
	})
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
		return nil, err
	}
	return client, nil
}
