package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Chat(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/v1")
	r.POST("/chat/completions", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.GetCompletions(db, c)
	})
	r.GET("/token_check", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.TokenCheck(db, c)
	})
	r.GET("/models", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.GetModels(db, c)
	})
}
