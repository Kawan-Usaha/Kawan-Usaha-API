package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	onceArticle     sync.Once
	articleInstance *ArticleSingleton
)

type ArticleSingleton struct {
	db *gorm.DB
	q  *gin.Engine
}

func GetArticleInstance() *ArticleSingleton {
	onceArticle.Do(func() {
		articleInstance = &ArticleSingleton{}
	})
	return articleInstance
}

func (a *ArticleSingleton) Init(db *gorm.DB, q *gin.Engine) {
	a.db = db
	a.q = q
}

func (a *ArticleSingleton) SetupRoutes() {
	r := a.q.Group("/article")
	r.Static("/images", "./images")
	// Get all articles
	r.GET("", func(c *gin.Context) {
		controller.ListAllArticles(a.db, c)
	})
	// Get article by id
	r.GET("/content", func(c *gin.Context) {
		controller.GetArticle(a.db, c)
	})
	r.POST("/favorite", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.AddToFavorites(a.db, c)
	})
	// Get article by user id
	r.GET("/owned", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.ListOwnedArticles(a.db, c)
	})
	// Search owned article by title
	r.GET("/owned/search", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.SearchOwnedArticles(a.db, c)
	})
	// Search all article by title
	r.GET("/search", func(c *gin.Context) {
		controller.SearchAllArticles(a.db, c)
	})
	// Search all article by title
	r.GET("/category", func(c *gin.Context) {
		controller.SearchArticlebyCategory(a.db, c)
	})
	// Create article
	r.POST("/create", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.CreateArticle(a.db, c)
	})
	// Update article
	r.PATCH("/update", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UpdateArticle(a.db, c)
	})
	// Delete article
	r.DELETE("/delete", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.DeleteArticle(a.db, c)
	})
}
