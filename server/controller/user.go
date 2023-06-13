package controller

import (
	"encoding/json"
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
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
	var err error
	input.Image, err = lib.Compare(updatedImage, user.Image, c.Request.Context())

	if err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to update user", err.Error()))
		return
	}

	if user.Email != input.Email {
		user.Verified = false
		user.Email = input.Email
	}
	user.Name = input.Name
	user.Image = input.Image

	if err := db.Save(&user).Error; err != nil {
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

func GetFavoriteArticles(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var user Model.User
	if err := db.Where("user_id = ?", subs).First(&user).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get user", err.Error()))
		return
	}
	var articles []Model.Article
	if err := db.Model(&user).Association("FavoriteArticles").Find(&articles); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get favorite articles", err.Error()))
		return
	}
	var articleData []gin.H
	for _, article := range articles {
		articleData = append(articleData, gin.H{
			"id":          article.ID,
			"title":       article.Title,
			"image":       article.Image,
			"isGenerated": article.IsGenerated,
			"createdAt":   article.CreatedAt,
			"updatedAt":   article.UpdatedAt,
		})
	}
	result := gin.H{
		"articles": articleData,
	}
	c.JSON(200, lib.OkResponse("Success get favorite articles", result))
}
