package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/user")
	// Get logged in User profile
	r.GET("/profile", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UserProfile(db, c)
	})
	// Update logged in User profile
	r.PATCH("/profile", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UpdateUserProfile(db, c)
	})
}
