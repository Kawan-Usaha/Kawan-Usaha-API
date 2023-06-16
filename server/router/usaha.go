package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	onceUsaha     sync.Once
	usahaInstance *UsahaSingleton
)

type UsahaSingleton struct {
	db *gorm.DB
	q  *gin.Engine
}

func GetUsahaInstance() *UsahaSingleton {
	onceUsaha.Do(func() {
		usahaInstance = &UsahaSingleton{}
	})
	return usahaInstance
}

func (u *UsahaSingleton) Init(db *gorm.DB, q *gin.Engine) {
	u.db = db
	u.q = q
}

func (u *UsahaSingleton) SetupRoutes() {
	r := u.q.Group("/usaha")

	r.GET("/list", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.ListOwnedUsaha(u.db, c)
	})
	r.GET("/search", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.SearchOwnedUsahaByTitle(u.db, c)
	})
	r.GET("/detail", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.GetOwnedUsahaByID(u.db, c)
	})
	r.POST("/create", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.CreateUsaha(u.db, c)
	})
	r.PATCH("/update", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UpdateUsaha(u.db, c)
	})
	r.DELETE("/delete", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.DeleteUsaha(u.db, c)
	})
}
