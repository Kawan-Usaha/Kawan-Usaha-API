package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Usaha(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/usaha")

	r.GET("/list", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.ListOwnedUsaha(db, c)
	})
	r.GET("/search", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.SearchOwnedUsahaByTitle(db, c)
	})
	r.GET("/detail", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.GetOwnedUsahaByID(db, c)
	})
	r.POST("/create", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.CreateUsaha(db, c)
	})
	r.DELETE("/delete", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.DeleteUsaha(db, c)
	})
}
