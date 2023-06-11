package controller

import (
	"encoding/json"
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListOwnedArticles(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)

	// Retrieve page and page size from query parameters
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1 // Set default page to 1 if invalid or not provided
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10 // Set default page size to 10 if invalid or not provided
	}

	var totalArticles int64
	if err := db.Model(&Model.Article{}).Where("user_id = ?", subs).Count(&totalArticles).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}

	var articles []Model.Article
	offset := (page - 1) * pageSize
	if err := db.Preload("Category").Where("user_id = ?", subs).Offset(offset).Limit(pageSize).Order("updated_at desc").Find(&articles).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}

	result := make([]gin.H, len(articles))
	for i, v := range articles {
		result[i] = gin.H{
			"id":           v.ID,
			"title":        v.Title,
			"is_published": v.IsPublished,
			"category":     v.Category,
			"created_at":   v.CreatedAt,
			"updated_at":   v.UpdatedAt,
			"user_id":      v.UserID,
			"image":        v.Image,
		}
	}

	response := gin.H{
		"total_articles": totalArticles,
		"page":           page,
		"page_size":      pageSize,
		"articles":       result,
	}

	c.JSON(200, lib.OkResponse("Success get article", response))
}

func ListAllArticles(db *gorm.DB, c *gin.Context) {
	// Retrieve page and page size from query parameters
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1 // Set default page to 1 if invalid or not provided
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10 // Set default page size to 10 if invalid or not provided
	}

	var totalArticles int64
	if err := db.Model(&Model.Article{}).Count(&totalArticles).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}

	var articles []Model.Article
	offset := (page - 1) * pageSize
	if err := db.Preload("Category").Offset(offset).Limit(pageSize).Order("updated_at desc").Find(&articles).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}

	result := make([]gin.H, len(articles))
	for i, v := range articles {
		result[i] = gin.H{
			"id":           v.ID,
			"title":        v.Title,
			"is_published": v.IsPublished,
			"category":     v.Category,
			"created_at":   v.CreatedAt,
			"updated_at":   v.UpdatedAt,
			"image":        v.Image,
		}
	}

	response := gin.H{
		"total_articles": totalArticles,
		"page":           page,
		"page_size":      pageSize,
		"articles":       result,
	}

	c.JSON(200, lib.OkResponse("Success get article", response))
}

func GetArticle(db *gorm.DB, c *gin.Context) {
	var article Model.Article
	if err := db.Where("id = ?", c.Query("id")).Preload("User").Preload("Category").First(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}
	response := gin.H{
		"id":           article.ID,
		"title":        article.Title,
		"content":      article.Content,
		"is_published": article.IsPublished,
		"category":     article.Category,
		"user":         article.User.Name,
		"created_at":   article.CreatedAt,
		"updated_at":   article.UpdatedAt,
		"image":        article.Image,
	}
	c.JSON(200, lib.OkResponse("Success get article", response))
}

func SearchOwnedArticles(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)

	// Retrieve page and page size from query parameters
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1 // Set default page to 1 if invalid or not provided
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10 // Set default page size to 10 if invalid or not provided
	}

	var totalArticles int64
	query := db.Model(&Model.Article{}).Where("user_id = ? AND LOWER(title) LIKE ?", subs, "%"+c.Query("title")+"%")
	if err := query.Count(&totalArticles).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}

	var articles []Model.Article
	offset := (page - 1) * pageSize
	if err := query.Preload("Category").Offset(offset).Limit(pageSize).Order("updated_at desc").Find(&articles).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}

	result := make([]gin.H, len(articles))
	for i, v := range articles {
		result[i] = gin.H{
			"id":           v.ID,
			"title":        v.Title,
			"is_published": v.IsPublished,
			"category":     v.Category,
			"created_at":   v.CreatedAt,
			"updated_at":   v.UpdatedAt,
			"image":        v.Image,
		}
	}

	response := gin.H{
		"total_articles": totalArticles,
		"page":           page,
		"page_size":      pageSize,
		"articles":       result,
	}

	c.JSON(200, lib.OkResponse("Success get article", response))
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
			"image":        v.Image,
		})
	}

	response := gin.H{
		"page":      page,
		"page_size": pageSize,
		"total":     total,
		"articles":  result,
	}

	c.JSON(200, lib.OkResponse("Success get articles", response))
}

func CreateArticle(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var requestData struct {
		Article  Model.Article `json:"article"`
		Category uint          `json:"category"`
	}
	article := c.PostForm("article")
	json.Unmarshal([]byte(article), &requestData)

	// Set the user ID
	requestData.Article.UserID = subs
	requestData.Article.CreatedAt = time.Now()
	requestData.Article.UpdatedAt = time.Now()

	// Find the category
	var category Model.Category
	if err := db.First(&category, requestData.Category).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to find category", err.Error()))
		return
	}

	image, _ := c.FormFile("image")
	if image != nil {
		var imagePath string
		var err error
		if os.Getenv("DEPLOYMENT_MODE") == "local" {
			imagePath, err = lib.SaveImageOffline(image, "/article")
		} else {
			imagePath, err = lib.SaveImageOnline(image)
		}
		if err != nil {
			c.JSON(500, lib.ErrorResponse("Failed to save image", err.Error()))
			return
		}
		requestData.Article.Image = imagePath
	} else {
		requestData.Article.Image = ""
	}
	// Assign the category to the article
	requestData.Article.Category = []Model.Category{category}

	if err := db.Create(&requestData.Article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to create article", err.Error()))
		return
	}

	result := gin.H{
		"id":           requestData.Article.ID,
		"title":        requestData.Article.Title,
		"is_published": requestData.Article.IsPublished,
		"category":     requestData.Article.Category,
		"created_at":   requestData.Article.CreatedAt,
		"updated_at":   requestData.Article.UpdatedAt,
		"image":        requestData.Article.Image,
	}

	c.JSON(200, lib.OkResponse("Success create article", result))
}

func UpdateArticle(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var input struct {
		Article  Model.Article `json:"article"`
		Category uint          `json:"category"`
	}
	articlein := c.PostForm("article")
	json.Unmarshal([]byte(articlein), &input)

	var article Model.Article
	if err := db.Where("id = ? AND user_id = ?", input.Article.ID, subs).First(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}

	// Update the article properties
	article.Title = input.Article.Title
	article.Content = input.Article.Content
	article.IsPublished = input.Article.IsPublished
	article.UpdatedAt = time.Now()

	updatedImage, _ := c.FormFile("image")
	if updatedImage != nil {
		// Calculate the MD5 hash of the updated image
		updatedHash, err := lib.CalculateMD5Hash(updatedImage)
		if err != nil {
			c.JSON(500, lib.ErrorResponse("Failed to calculate hash for the updated image", err.Error()))
			return
		}

		var existingHash string
		if article.Image != "" {
			// Calculate the MD5 hash of the existing image
			if os.Getenv("DEPLOYMENT_MODE") == "local" {
				existingHash, err = lib.CalculateMD5HashFromOffline(article.Image)
				if err != nil {
					c.JSON(500, lib.ErrorResponse("Failed to calculate hash for the existing image", err.Error()))
					return
				}
			} else {
				existingHash, err = lib.CalculateMD5HashFromURL(article.Image)
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
				if article.Image != "" {
					if err := lib.DeleteImageOffline(article.Image); err != nil {
						c.JSON(500, lib.ErrorResponse("Failed to delete image", err.Error()))
						return
					}
				}
				imagePath, err = lib.SaveImageOffline(updatedImage, "/article")
			} else {
				if article.Image != "" {
					if err := lib.DeleteImageOnline(article.Image); err != nil {
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
			article.Image = imagePath
			log.Println("Updated image")
		} else {
			log.Println("Image not updated")
		}
	} else {
		if os.Getenv("DEPLOYMENT_MODE") == "local" {
			if article.Image != "" {
				if err := lib.DeleteImageOffline(article.Image); err != nil {
					c.JSON(500, lib.ErrorResponse("Failed to delete image", err.Error()))
					return
				}
			}
		} else {
			if article.Image != "" {
				if err := lib.DeleteImageOnline(article.Image); err != nil {
					c.JSON(500, lib.ErrorResponse("Failed to delete image", err.Error()))
					return
				}
			}
		}
		article.Image = ""
	}

	// Find the category
	var category Model.Category
	if err := db.First(&category, input.Category).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to find category", err.Error()))
		return
	}

	// Assign the category to the article
	article.Category = []Model.Category{category}

	if err := db.Save(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to update article", err.Error()))
		return
	}

	result := gin.H{
		"id":           article.ID,
		"user_id":      article.UserID,
		"title":        article.Title,
		"is_published": article.IsPublished,
		"category":     article.Category,
		"content":      article.Content,
		"created_at":   article.CreatedAt,
		"updated_at":   article.UpdatedAt,
		"image":        article.Image,
	}

	c.JSON(200, lib.OkResponse("Success update article", result))
}

func DeleteArticle(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var del struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&del); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to bind json", err.Error()))
		return
	}
	var article Model.Article
	if err := db.Preload("Category").Preload("User").Where("id = ? AND user_id = ?", del.ID, subs).First(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get article", err.Error()))
		return
	}

	// Clear the article's associations with categories and users
	db.Model(&article).Association("Category").Clear()
	db.Model(&article).Association("User").Clear()

	if err := db.Delete(&article).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to delete article", err.Error()))
		return
	}
	if os.Getenv("DEPLOYMENT_MODE") == "local" {
		if article.Image != "" {
			if err := lib.DeleteImageOffline(article.Image); err != nil {
				c.JSON(500, lib.ErrorResponse("Failed to delete image", err.Error()))
				return
			}
		}
	} else {
		if article.Image != "" {
			if err := lib.DeleteImageOnline(article.Image); err != nil {
				c.JSON(500, lib.ErrorResponse("Failed to delete image", err.Error()))
				return
			}
		}
	}
	c.JSON(200, lib.OkResponse("Success delete article", nil))
}
