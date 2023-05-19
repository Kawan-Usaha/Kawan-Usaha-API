package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/user")
	r.GET("", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UserProfile(db, c)
	})
}
