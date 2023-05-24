package controller

import (
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"strconv"
	"time"

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

func GetArticle(db *gorm.DB, c *gin.Context) {
	var article Model.Article
	if err := db.Where("id = ?", c.Query("id")).First(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}
	c.JSON(200, lib.OkResponse("Success get article", article))
}

func SearchOwnedArticles(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var article []Model.Article
	if err := db.Where("user_id = ? AND LOWER(title) LIKE ?", subs, "%"+c.Query("title")+"%").Find(&article).Error; err != nil {
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

func SearchAllArticles(db *gorm.DB, c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		pageSize = 10
	}

	var total int64
	if err := db.Model(&Model.Article{}).Where("LOWER(title) LIKE ?", "%"+c.Query("title")+"%").Count(&total).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to count articles", err.Error()))
		return
	}

	var articles []Model.Article
	if err := db.Where("LOWER(title) LIKE ?", "%"+c.Query("title")+"%").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("updated_at desc"). // Sort by updated_at field in descending order
		Find(&articles).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get articles", err.Error()))
		return
	}

	result := make([]gin.H, 0)
	for _, v := range articles {
		result = append(result, gin.H{
			"id":           v.ID,
			"title":        v.Title,
			"is_published": v.IsPublished,
			"category":     v.Category,
			"created_at":   v.CreatedAt,
			"updated_at":   v.UpdatedAt,
		})
	}

	response := gin.H{
		"page":     page,
		"pageSize": pageSize,
		"total":    total,
		"articles": result,
	}

	c.JSON(200, lib.OkResponse("Success get articles", response))
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
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	if err := db.Create(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to create article", err.Error()))
		return
	}
	result := gin.H{
		"id":           article.ID,
		"title":        article.Title,
		"is_published": article.IsPublished,
		"category":     article.Category,
		"created_at":   article.CreatedAt,
		"updated_at":   article.UpdatedAt,
	}
	c.JSON(200, lib.OkResponse("Success create article", result))
}

func UpdateArticle(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var input Model.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to bind json", err.Error()))
		return
	}
	var article Model.Article
	if err := db.Where("id = ? AND user_id = ?", input.ID, subs).First(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}
	article.UpdatedAt = time.Now()
	if err := db.Model(&article).Updates(input).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to update article", err.Error()))
		return
	}

	result := gin.H{
		"id":           article.ID,
		"user_id":      article.UserID, // "user_id" is not in the model, but it is in the response
		"title":        article.Title,
		"is_published": article.IsPublished,
		"category":     article.Category,
		"content":      article.Content,
		"created_at":   article.CreatedAt,
		"updated_at":   article.UpdatedAt,
	}
	c.JSON(200, lib.OkResponse("Success update article", result))
}

func DeleteArticle(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var del delete
	if err := c.ShouldBindJSON(&del); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to bind json", err.Error()))
		return
	}
	var article Model.Article
	if err := db.Where("id = ? AND user_id = ?", del.ID, subs).First(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}
	db.Model(&article).Association("Category").Clear()
	db.Model(&article).Association("User").Clear()
	if err := db.Delete(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to delete article", err.Error()))
		return
	}
	c.JSON(200, lib.OkResponse("Success delete article", nil))
}
