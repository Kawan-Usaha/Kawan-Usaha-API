package router

// category.go

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	onceCategory     sync.Once
	categoryInstance *CategorySingleton
)

type CategorySingleton struct {
	db *gorm.DB
	q  *gin.Engine
}

func GetCategoryInstance() *CategorySingleton {
	onceCategory.Do(func() {
		categoryInstance = &CategorySingleton{}
	})
	return categoryInstance
}

func (c *CategorySingleton) Init(db *gorm.DB, q *gin.Engine) {
	c.db = db
	c.q = q
}

func (cat *CategorySingleton) SetupRoutes() {
	r := cat.q.Group("/category")
	r.GET("", func(c *gin.Context) {
		controller.ListAllCategory(cat.db, c)
	})
	r.GET("/detail", func(c *gin.Context) {
		controller.GetCategory(cat.db, c)
	})
	r.GET("/search", func(c *gin.Context) {
		controller.SearchCategoryByName(cat.db, c)
	})
	r.POST("/create", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.CreateCategory(cat.db, c)
	})
	r.PATCH("/update", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UpdateCategory(cat.db, c)
	})
	r.DELETE("/delete", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.DeleteCategory(cat.db, c)
	})
}
