package controller

import (
	"encoding/json"
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserProfile(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var user Model.User
	if err := db.Where("user_id = ?", subs).First(&user).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get user", err.Error()))
		return
	}
	result := gin.H{
		"userId":    user.UserId,
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
		"verified":  user.Verified,
		"usaha":     user.Usaha,
		"articles":  user.Article,
		"image":     user.Image,
	}
	c.JSON(200, lib.OkResponse("Success get user", result))
}

func UpdateUserProfile(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var input Model.User

	request := c.PostForm("user")
	json.Unmarshal([]byte(request), &input)

	input.UpdatedAt = time.Now()
	var user Model.User
	if err := db.Where("user_id = ?", subs).First(&user).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get user", err.Error()))
		return
	}
	updatedImage, _ := c.FormFile("image")
	if updatedImage != nil {
		// Calculate the MD5 hash of the updated image
		updatedHash, err := lib.CalculateMD5Hash(updatedImage)
		if err != nil {
			c.JSON(500, lib.ErrorResponse("Failed to calculate hash for the updated image", err.Error()))
			return
		}

		var existingHash string
		if user.Image != "" {
			// Calculate the MD5 hash of the existing image
			if os.Getenv("DEPLOYMENT_MODE") == "local" {
				existingHash, err = lib.CalculateMD5HashFromOffline(user.Image)
				if err != nil {
					c.JSON(500, lib.ErrorResponse("Failed to calculate hash for the existing image", err.Error()))
					return
				}
			} else {
				existingHash, err = lib.CalculateMD5HashFromURL(user.Image)
				if err != nil {
					c.JSON(500, lib.ErrorResponse("Failed to calculate hash for the existing image", err.Error()))
					return
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
				if user.Image != "" {
					if err := lib.DeleteImageOffline(user.Image); err != nil {
						c.JSON(500, lib.ErrorResponse("Failed to delete image", err.Error()))
						return
					}
				}
				imagePath, err = lib.SaveImageOffline(updatedImage, "/article")
			} else {
				if user.Image != "" {
					if err := lib.DeleteImageOnline(user.Image); err != nil {
						c.JSON(500, lib.ErrorResponse("Failed to delete image", err.Error()))
						return
					}
				}
				imagePath, err = lib.SaveImageOnline(updatedImage)
			}
			if err != nil {
				c.JSON(500, lib.ErrorResponse("Failed to save image", err.Error()))
				return
			}
			input.Image = imagePath
			log.Println("Updated image")
		} else {
			log.Println("Image not updated")
		}
	} else {
		if os.Getenv("DEPLOYMENT_MODE") == "local" {
			if user.Image != "" {
				if err := lib.DeleteImageOffline(user.Image); err != nil {
					c.JSON(500, lib.ErrorResponse("Failed to delete image", err.Error()))
					return
				}
			}
		} else {
			if user.Image != "" {
				if err := lib.DeleteImageOnline(user.Image); err != nil {
					c.JSON(500, lib.ErrorResponse("Failed to delete image", err.Error()))
					return
				}
			}
		}
		input.Image = ""
	}
	if err := db.Model(&user).Updates(input).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to update user", err.Error()))
		return
	}
	result := gin.H{
		"name":     user.Name,
		"email":    user.Email,
		"verified": user.Verified,
		"image":    user.Image,
	}
	c.JSON(200, lib.OkResponse("Success update user", result))
}
