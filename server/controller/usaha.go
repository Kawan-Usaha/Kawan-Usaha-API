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
	if err := db.Preload("Tags").Where("user_id = ? AND LOWER(usaha_name) LIKE ?", subs, "%"+strings.ToLower(c.Query("name"))+"%").Find(&usaha).Error; err != nil {
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

func UpdateUsaha(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var input Model.Usaha
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to bind JSON", err.Error()))
		return
	}
	var usaha Model.Usaha
	if err := db.Preload("User").Preload("Tags").Where("user_id = ?", subs).Where("id = ?", input.ID).First(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get usaha", err.Error()))
		return
	}
	input.UpdatedAt = time.Now()

	existingTagsMap := make(map[string]Model.Tag)
	updatedTags := make([]Model.Tag, 0)

	existingTags := make([]Model.Tag, 0)
	if err := db.Find(&existingTags).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get existing tags", err.Error()))
		return
	}
	for _, existingTag := range existingTags {
		existingTagsMap[strings.ToLower(existingTag.Name)] = existingTag
	}

	for _, inputTag := range input.Tags {
		tagName := strings.ToLower(inputTag.Name)
		existingTag, exists := existingTagsMap[tagName]

		if exists {
			updatedTags = append(updatedTags, existingTag)
		} else {
			newTag := Model.Tag{
				Name:      tagName,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.Create(&newTag).Error; err != nil {
				c.JSON(400, lib.ErrorResponse("Failed to create tag", err.Error()))
				return
			}
			updatedTags = append(updatedTags, newTag)
		}
	}

	usaha.Tags = updatedTags

	disconnectedTags := make([]Model.Tag, 0)
	if err := db.Model(&usaha).Association("Tags").Find(&disconnectedTags); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to find disconnected tags", err.Error()))
		return
	}

	tagsToDisconnect := make([]Model.Tag, 0)
	for _, tag := range disconnectedTags {
		found := false
		for _, inputTag := range input.Tags {
			if strings.EqualFold(strings.ToLower(inputTag.Name), strings.ToLower(tag.Name)) {
				found = true
				break
			}
		}
		if !found {
			tagsToDisconnect = append(tagsToDisconnect, tag)
		}
	}

	if len(tagsToDisconnect) > 0 {
		if err := db.Model(&usaha).Association("Tags").Delete(tagsToDisconnect); err != nil {
			c.JSON(400, lib.ErrorResponse("Failed to disconnect tags from Usaha", err.Error()))
			return
		}
	}

	if err := db.Save(&usaha).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to update usaha", err.Error()))
		return
	}
	if err := db.Preload("User").Preload("Tags").Where("user_id = ?", subs).Where("id = ?", input.ID).First(&usaha).Error; err != nil {
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
	c.JSON(200, lib.OkResponse("Success update usaha", result))
}

type deleteUsaha struct {
	ID int `json:"id"`
}

func DeleteUsaha(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)

	var del deleteUsaha
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
