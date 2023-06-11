package lib

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DeleteImageOnline(imageURL string) error {
	client, err := GetStorageClient()
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
		return err
	}

	ctx := context.Background()

	// Extract the bucket and object name from the image URL
	bucketName, objectName, err := extractBucketAndObjectName(imageURL)
	if err != nil {
		log.Panicf("Failed to extract bucket and object name: %v", err)
		return err
	}

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(objectName)

	if err := obj.Delete(ctx); err != nil {
		log.Panicf("Failed to delete object: %v", err)
		return err
	}

	return nil
}

func extractBucketAndObjectName(imageURL string) (string, string, error) {
	// Find the index of "/b/" and "/o/"
	bIndex := strings.Index(imageURL, "/b/")
	oIndex := strings.Index(imageURL, "/o/")

	if bIndex == -1 || oIndex == -1 {
		return "", "", fmt.Errorf("invalid image URL")
	}

	// Extract the bucket name
	bucketStart := bIndex + 3
	bucketEnd := oIndex
	bucketName := imageURL[bucketStart:bucketEnd]

	// Extract the object name
	objectStart := oIndex + 3
	objectEnd := strings.Index(imageURL[objectStart:], "?")
	if objectEnd == -1 {
		objectEnd = len(imageURL)
	} else {
		objectEnd += objectStart
	}
	objectName := imageURL[objectStart:objectEnd]

	return bucketName, objectName, nil
}

func DeleteImageOffline(imagePath string) error {
	// Extract the file name from the image path
	fileName := extractFileNameFromImagePath(imagePath)

	// Construct the file path
	filePath := filepath.Join("images", fileName)

	// Delete the file
	err := os.Remove(filePath)
	if err != nil {
		log.Panicf("Failed to delete image file: %v", err)
		return err
	}

	return nil
}

func extractFileNameFromImagePath(imagePath string) string {
	// Remove the custom port number, if present
	imagePath = strings.SplitN(imagePath, "/", 3)[2]

	// Find the last occurrence of '/'
	lastIndex := strings.LastIndex(imagePath, "/")
	if lastIndex == -1 {
		return ""
	}

	// Extract the file name
	fileName := imagePath[lastIndex+1:]

	return fileName
}
