package lib

import (
	"context"
	"log"
	"mime/multipart"
	"os"
)

func Compare(updatedImage *multipart.FileHeader, oldImage string, ctx context.Context) (string, error) {
	image := ""
	if updatedImage != nil {
		// Calculate the MD5 hash of the updated image
		updatedHash, err := CalculateMD5Hash(updatedImage)
		if err != nil {
			return "", err
		}

		var existingHash string
		if oldImage != "" {
			// Calculate the MD5 hash of the existing image
			if os.Getenv("DEPLOYMENT_MODE") == "local" {
				existingHash, err = CalculateMD5HashFromOffline(oldImage)
				if err != nil {
					return "", err
				}
			} else {
				existingHash, err = CalculateMD5HashFromURL(oldImage)
				if err != nil {
					return "", err
				}
			}
		} else {
			existingHash = ""
		}
		// Compare the hashes to determine if the images are identical
		if updatedHash != existingHash {
			// Updated image is different, overwrite the existing image
			var imagePath string
			var err error
			if os.Getenv("DEPLOYMENT_MODE") == "local" {
				if oldImage != "" {
					if err := DeleteImageOffline(oldImage); err != nil {
						return "", err
					}
				}
				imagePath, err = SaveImageOffline(updatedImage, "/article")
			} else {
				if oldImage != "" {
					if err := DeleteImageOnline(oldImage); err != nil {
						return "", err
					}
				}
				imagePath, err = SaveImageOnline(updatedImage)
			}
			if err != nil {
				return "", err
			}
			image = imagePath
			log.Println("Updated image")
		} else {
			log.Println("Image not updated")
		}
	} else {
		if os.Getenv("DEPLOYMENT_MODE") == "local" {
			if oldImage != "" {
				if err := DeleteImageOffline(oldImage); err != nil {
					return "", err
				}
			}
		} else {
			if oldImage != "" {
				if err := DeleteImageOnline(oldImage); err != nil {
					return "", err
				}
			}
		}
		image = ""
	}
	return image, nil
}
