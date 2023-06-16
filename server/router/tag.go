package router

import (
	"kawan-usaha-api/server/controller"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	onceTag     sync.Once
	tagInstance *TagSingleton
)

type TagSingleton struct {
	db *gorm.DB
	q  *gin.Engine
}

func GetTagInstance() *TagSingleton {
	onceTag.Do(func() {
		tagInstance = &TagSingleton{}
	})
	return tagInstance
}

func (a *TagSingleton) Init(db *gorm.DB, q *gin.Engine) {
	a.db = db
	a.q = q
}

func (a *TagSingleton) SetupRoutes() {
	r := a.q.Group("/tag")
	// Get all tags
	r.GET("", func(c *gin.Context) {
		controller.ListAllTags(a.db, c)
	})
}
