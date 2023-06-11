package lib

import (
	"math/rand"
	"time"
)

func generateRandomFileName(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Define a list of characters that can be used in the file name
	// Modify this according to your requirements
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	fileName := make([]byte, length)
	for i := 0; i < length; i++ {
		fileName[i] = characters[rand.Intn(len(characters))]
	}

	return string(fileName)
}
