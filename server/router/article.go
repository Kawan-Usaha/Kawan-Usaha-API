package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Article(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/article")
	// Get all articles
	r.GET("/all", func(c *gin.Context) {
		controller.ListAllArticles(db, c)
	})
	// Get article by id
	r.GET("", func(c *gin.Context) {
		controller.GetArticle(db, c)
	})
	// Get article by user id
	r.GET("/owned", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.ListOwnedArticles(db, c)
	})
	// Search owned article by title
	r.GET("/owned/search", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.SearchOwnedArticles(db, c)
	})
	// Search all article by title
	r.GET("/search", func(c *gin.Context) {
		controller.SearchAllArticles(db, c)
	})
	// Create article
	r.POST("/create", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.CreateArticle(db, c)
	})
	// Update article
	r.PATCH("/update", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UpdateArticle(db, c)
	})
	// Delete article
	r.DELETE("/delete", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.DeleteArticle(db, c)
	})
}
