package controller

import (
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"

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
	}
	c.JSON(200, lib.OkResponse("Success get user", result))
}
