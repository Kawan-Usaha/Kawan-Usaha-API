package controller

import (
	"errors"
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListOwnedUsaha(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var usaha []Model.Usaha
	if err := db.Preload("User").Preload("Tags").Where("user_id = ?", subs).Find(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get usaha", err.Error()))
		return
	}
	result := make([]gin.H, 0)
	for _, v := range usaha {
		result = append(result, gin.H{
			"id":         v.ID,
			"user_id":    v.UserID,
			"usaha_name": v.UsahaName,
			"type":       v.Type,
			"tags":       v.Tags,
			"created_at": v.CreatedAt,
			"updated_at": v.UpdatedAt,
		})
	}
	c.JSON(200, lib.OkResponse("Success get usaha", result))
}

func SearchOwnedUsahaByTitle(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var usaha []Model.Usaha
	if err := db.Where("user_id = ? AND LOWER(usaha_name) LIKE ?", subs, "%"+strings.ToLower(c.Query("name"))+"%").Find(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get usaha", err.Error()))
		return
	}
	for_each := make([]gin.H, 0)
	for _, v := range usaha {
		for_each = append(for_each, gin.H{
			"id":         v.ID,
			"user_id":    v.UserID,
			"usaha_name": v.UsahaName,
			"type":       v.Type,
			"tags":       v.Tags,
			"created_at": v.CreatedAt,
			"updated_at": v.UpdatedAt,
		})
	}
	result := gin.H{"data": for_each}
	c.JSON(200, lib.OkResponse("Success get usaha", result))
}

func GetOwnedUsahaByID(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var usaha Model.Usaha
	if err := db.Preload("tags").Where("user = ? AND id = ?", subs, c.Param("id")).First(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get usaha", err.Error()))
		return
	}
	c.JSON(200, lib.OkResponse("Success get usaha", usaha))
}

func CreateUsaha(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)

	var usaha Model.Usaha
	if err := c.ShouldBindJSON(&usaha); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to bind JSON", err.Error()))
		return
	}

	usaha.UserID = subs
	usaha.CreatedAt = time.Now()
	usaha.UpdatedAt = time.Now()

	for i := range usaha.Tags {
		tagName := strings.ToLower(usaha.Tags[i].Name)
		var existingTag Model.Tag
		err := db.Where("LOWER(name) = ?", tagName).First(&existingTag).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			newTag := Model.Tag{
				Name:      tagName,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.Create(&newTag).Error; err != nil {
				c.JSON(400, lib.ErrorResponse("Failed to create tag", err.Error()))
				return
			}
			usaha.Tags[i] = newTag
		} else if err != nil {
			c.JSON(400, lib.ErrorResponse("Failed to check existing tag", err.Error()))
			return
		} else {
			usaha.Tags[i] = existingTag
		}
	}

	if err := db.Create(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to create usaha", err.Error()))
		return
	}

	result := gin.H{
		"id":         usaha.ID,
		"user_id":    usaha.UserID,
		"usaha_name": usaha.UsahaName,
		"type":       usaha.Type,
		"tags":       usaha.Tags,
		"created_at": usaha.CreatedAt,
		"updated_at": usaha.UpdatedAt,
	}
	c.JSON(200, lib.OkResponse("Success create usaha", result))
}
