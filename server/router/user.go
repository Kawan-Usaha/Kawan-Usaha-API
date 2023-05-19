package router

import (
	"kawan-usaha-api/server/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/user")
	r.GET("", func(c *gin.Context) {
		c.JSON(200, lib.OkResponse("Welcome to Kawan Usaha API!", nil))
	})
}
