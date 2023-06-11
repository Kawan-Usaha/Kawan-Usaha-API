package lib

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func CalculateMD5Hash(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Helper function to calculate the MD5 hash of a file from its URL
func CalculateMD5HashFromURL(url string) (string, error) {
	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Calculate the MD5 hash of the downloaded file
	hash := md5.New()
	log.Println(resp.Body)
	if _, err := io.Copy(hash, resp.Body); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func CalculateMD5HashFromOffline(file string) (string, error) {
	f, err := os.Open(filepath.Join("images", string(extractFileNameFromImagePath(file))))
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	log.Println(f)
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
