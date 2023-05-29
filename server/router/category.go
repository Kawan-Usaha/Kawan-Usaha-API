package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Category(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/category")
	r.GET("", func(c *gin.Context) {
		controller.ListAllCategory(db, c)
	})
	r.GET("/detail", func(c *gin.Context) {
		controller.GetCategory(db, c)
	})
	r.GET("/search", func(c *gin.Context) {
		controller.SearchCategoryByName(db, c)
	})
	r.POST("/create", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.CreateCategory(db, c)
	})
	r.PATCH("/update", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UpdateCategory(db, c)
	})
	r.DELETE("/delete", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.DeleteCategory(db, c)
	})

}
