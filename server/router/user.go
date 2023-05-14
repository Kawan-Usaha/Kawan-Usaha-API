package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kawan-usaha-api/server/lib"
)

func User(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/user")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, lib.Ok("Welcome to Kawan Usaha API!", nil))
	})
}
