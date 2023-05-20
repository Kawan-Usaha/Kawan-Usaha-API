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
	if err := db.Preload("Tags").Where("user_id = ?", subs).Find(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get usaha", err.Error()))
		return
	}
	result := make([]gin.H, 0)
	for _, v := range usaha {
		result = append(result, gin.H{
			"id":         v.ID,
			"usaha_name": v.UsahaName,
			"type":       v.Type,
			"tags":       v.Tags,
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
			"usaha_name": v.UsahaName,
			"type":       v.Type,
			"tags":       v.Tags,
		})
	}
	result := gin.H{"data": for_each}
	c.JSON(200, lib.OkResponse("Success get usaha", result))
}

func GetOwnedUsahaByID(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var usaha Model.Usaha
	if err := db.Preload("User").Preload("Tags").Where("user_id = ?", subs).Where("id = ?", c.Query("id")).First(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get usaha", err.Error()))
		return
	}
	result := gin.H{
		"id":      usaha.ID,
		"user_id": usaha.UserID,
		"user": gin.H{
			"id":         usaha.User.ID,
			"user_id":    usaha.User.UserId,
			"name":       usaha.User.Name,
			"verified":   usaha.User.Verified,
			"role_id":    usaha.User.RoleId,
			"created_at": usaha.User.CreatedAt,
			"updated_at": usaha.User.UpdatedAt,
		},
		"usaha_name": usaha.UsahaName,
		"type":       usaha.Type,
		"tags":       usaha.Tags,
		"created_at": usaha.CreatedAt,
		"updated_at": usaha.UpdatedAt,
	}
	c.JSON(200, lib.OkResponse("Success get usaha", result))
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

type delete struct {
	ID int `json:"id"`
}

func DeleteUsaha(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)

	var del delete
	if err := c.ShouldBindJSON(&del); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to bind JSON", err.Error()))
		return
	}

	var usaha Model.Usaha
	if err := db.Where("user_id = ?", subs).Where("id = ?", del.ID).First(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get usaha", err.Error()))
		return
	}

	if err := db.Delete(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to delete usaha", err.Error()))
		return
	}

	c.JSON(200, lib.OkResponse("Success delete usaha", nil))
}
