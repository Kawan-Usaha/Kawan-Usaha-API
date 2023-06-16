package lib

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func createStorageClient() (*storage.Client, error) {
	ctx := context.Background()
	// Retrieve the path to the service account key file from the environment variable
	serviceAccountKeyPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	// Create the client with the specified service account key file
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	return client, nil
}

func GetStorageClient() (*storage.Client, error) {
	client, err := createStorageClient()
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
		return nil, err
	}
	return client, nil
}
