package router

import (
	"kawan-usaha-api/server/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/auth")
	r.POST("/register", func(c *gin.Context) {
		controller.Register(db, c)
	})
	r.POST("/login", func(c *gin.Context) {
		controller.Login(db, c)
	})
}
