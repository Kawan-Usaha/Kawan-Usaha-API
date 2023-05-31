package controller

import (
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAllCategory(db *gorm.DB, c *gin.Context) {
	var category []Model.Category
	if err := db.Find(&category).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get category", err.Error()))
		return
	}
	result := make([]gin.H, 0)
	for _, v := range category {
		result = append(result, gin.H{
			"id":         v.ID,
			"title":      v.Title,
			"image":      v.Image,
			"created_at": v.CreatedAt,
			"updated_at": v.UpdatedAt,
		})
	}
	c.JSON(200, lib.OkResponse("Success get category", result))
}

func GetCategory(db *gorm.DB, c *gin.Context) {
	var category Model.Category
	if err := db.Preload("Tags").Where("id = ?", c.Query("id")).First(&category).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get category", err.Error()))
		return
	}

	response := gin.H{
		"id":       category.ID,
		"title":    category.Title,
		"image":    category.Image,
		"tags":     category.Tags,
		"articles": category.Articles,
	}

	c.JSON(200, lib.OkResponse("Success get category", response))
}

func SearchCategoryByName(db *gorm.DB, c *gin.Context) {
	var category []Model.Category
	if err := db.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(c.Query("name"))+"%").Find(&category).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get category", err.Error()))
		return
	}

	if len(category) == 0 {
		c.JSON(404, lib.ErrorResponse("No category found", nil))
		return
	}

	result := make([]gin.H, 0)
	for _, v := range category {
		result = append(result, gin.H{
			"id":         v.ID,
			"title":      v.Title,
			"image":      v.Image,
			"created_at": v.CreatedAt,
			"updated_at": v.UpdatedAt,
		})
	}
	c.JSON(200, lib.OkResponse("Success get category", result))
}

func CreateCategory(db *gorm.DB, c *gin.Context) {
	var input Model.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to create category", err.Error()))
		return
	}
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	// Save the category in the database
	if err := db.Create(&input).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to create category", err.Error()))
		return
	}

	tagsIDs := c.QueryArray("tags") // Assuming the tag IDs are provided as query parameters

	// Associate tags with the created category
	if len(tagsIDs) > 0 {
		var tags []Model.Tag
		if err := db.Find(&tags, tagsIDs).Error; err != nil {
			c.JSON(400, lib.ErrorResponse("Failed to find tags", err.Error()))
			return
		}

		// Update the association between category and tags
		if err := db.Model(&input).Association("Tags").Append(&tags); err != nil {
			c.JSON(400, lib.ErrorResponse("Failed to associate tags with category", err.Error()))
			return
		}
	}

	result := gin.H{
		"id":         input.ID,
		"title":      input.Title,
		"image":      input.Image,
		"tags":       input.Tags,
		"created_at": input.CreatedAt,
		"updated_at": input.UpdatedAt,
	}
	c.JSON(200, lib.OkResponse("Success create category", result))
}

func UpdateCategory(db *gorm.DB, c *gin.Context) {
	var input Model.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to update category", err.Error()))
		return
	}

	// Retrieve the existing category from the database
	var category Model.Category
	if err := db.Preload("Tags").First(&category, input.ID).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get category", err.Error()))
		return
	}

	// Update the category's properties based on the input JSON
	category.Title = input.Title
	category.Image = input.Image
	category.UpdatedAt = time.Now()

	// Clear the existing tags associated with the category
	if err := db.Model(&category).Association("Tags").Clear(); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to update category", err.Error()))
		return
	}

	// Add the new tags to the category
	for _, tag := range input.Tags {
		newTag := Model.Tag{Name: tag.Name}
		if err := db.FirstOrCreate(&newTag, &newTag).Error; err != nil {
			c.JSON(400, lib.ErrorResponse("Failed to update category", err.Error()))
			return
		}
		category.Tags = append(category.Tags, newTag)
	}

	// Save the updated category with its associated tags
	if err := db.Save(&category).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to update category", err.Error()))
		return
	}

	c.JSON(200, lib.OkResponse("Success update category", nil))
}

func DeleteCategory(db *gorm.DB, c *gin.Context) {
	var input Model.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to delete category", err.Error()))
		return
	}

	// Retrieve the existing category from the database
	var existingCategory Model.Category
	if err := db.Where("id = ?", input.ID).Preload("Tags").First(&existingCategory).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get category", err.Error()))
		return
	}

	// Clear the association between category and tags
	if err := db.Model(&existingCategory).Association("Tags").Clear(); err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to clear tags association", err.Error()))
		return
	}

	// Delete the category
	if err := db.Delete(&existingCategory).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to delete category", err.Error()))
		return
	}

	c.JSON(200, lib.OkResponse("Success delete category", nil))
}
