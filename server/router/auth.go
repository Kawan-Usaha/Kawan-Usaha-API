package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"

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
	r.POST("/forgot-password", func(c *gin.Context) {
		controller.EmailVerificationCodeFromForgotPassword(db, c)
	})
	r.POST("verify", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.EmailVerificationNormal(db, c)
	})
	r.POST("/generate", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.EmailVerificationCodeFromProfile(db, c)
	})
	r.POST("/forgot-password/generate", func(c *gin.Context) {
		controller.EmailVerificationCodeFromForgotPassword(db, c)
	})
	r.POST("/forgot-password/verify", func(c *gin.Context) {
		controller.EmailVerificationForgotPassword(db, c)
	})
}
