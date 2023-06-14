package controller

import (
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAllTags(db *gorm.DB, c *gin.Context) {
	// Retrieve page and page size from query parameters
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1 // Set default page to 1 if invalid or not provided
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10 // Set default page size to 10 if invalid or not provided
	}

	var totalTags int64
	if err := db.Model(&Model.Tag{}).Count(&totalTags).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get tags", err.Error()))
		return
	}

	var tags []Model.Tag
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Order("updated_at desc").Find(&tags).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get tags", err.Error()))
		return
	}

	result := make([]gin.H, len(tags))
	for i, v := range tags {
		result[i] = gin.H{
			"id":         v.ID,
			"name":       v.Name,
			"created_at": v.CreatedAt,
			"updated_at": v.UpdatedAt,
		}
	}
	response := gin.H{
		"total_tags": totalTags,
		"page":       page,
		"page_size":  pageSize,
		"tag":        result,
	}
	c.JSON(200, lib.OkResponse("Success get tags", response))
}
