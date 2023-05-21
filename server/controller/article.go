package controller

import (
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListOwnedArticles(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var article []Model.Article
	if err := db.Preload("Category").Where("user_id = ?", subs).Find(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}
	result := make([]gin.H, 0)
	for _, v := range article {
		result = append(result, gin.H{
			"id":           v.ID,
			"title":        v.Title,
			"is_published": v.IsPublished,
			"category":     v.Category,
			"created_at":   v.CreatedAt,
			"updated_at":   v.UpdatedAt,
		})
	}
	c.JSON(200, lib.OkResponse("Success get article", result))
}

func ListAllArticles(db *gorm.DB, c *gin.Context) {
	var article []Model.Article
	if err := db.Preload("Category").Find(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}
	result := make([]gin.H, 0)
	for _, v := range article {
		result = append(result, gin.H{
			"id":           v.ID,
			"title":        v.Title,
			"is_published": v.IsPublished,
			"category":     v.Category,
			"created_at":   v.CreatedAt,
			"updated_at":   v.UpdatedAt,
		})
	}
	c.JSON(200, lib.OkResponse("Success get article", result))
}

func CreateArticle(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var article Model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to bind json", err.Error()))
		return
	}
	article.UserID = subs
	if err := db.Create(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to create article", err.Error()))
		return
	}
	c.JSON(200, lib.OkResponse("Success create article", article))
}
