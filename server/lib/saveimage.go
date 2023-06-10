package lib

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveImageOffline(fileHeader *multipart.FileHeader, api_path string) (string, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create the destination file
	var filename string
	for {
		filename = generateRandomFileName(16) + filepath.Ext(fileHeader.Filename)
		if _, err := os.Stat(filepath.Join("images", string(filename))); os.IsNotExist(err) {
			break
		}
	}
	dst, err := os.Create(filepath.Join("images", string(filename)))
	if err != nil {
		return "file exist", err
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	baseurl := os.Getenv("BASE_URL")

	// Return the path to the saved image
	imagePath := filepath.Join(baseurl+api_path+"/images", filename)
	return imagePath, nil
}

func SaveImageOnline(fileheader *multipart.FileHeader) (string, error) {
	file, err := fileheader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	filename := generateRandomFileName(32) + filepath.Ext(fileheader.Filename)

	client, err := GetStorageClient()
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
		return "", err
	}

	ctx := context.Background()
	bucket := client.Bucket(os.Getenv("GOOGLE_BUCKET_NAME"))
	obj := bucket.Object(filename)

	wc := obj.NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		wc.Close()
		log.Panicf("Failed to write object: %v", err)
		return "", err
	}

	if err := wc.Close(); err != nil {
		log.Panicf("Failed to close writer: %v", err)
		return "", err
	}

	attrs, err := obj.Attrs(ctx)

	if err != nil {
		log.Panicf("Failed to get object attrs: %v", err)
		return "", err
	}
	file.Seek(0, io.SeekStart)
	client.Close()
	return attrs.MediaLink, nil
}
