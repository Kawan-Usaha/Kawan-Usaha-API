package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/auth")
	// Register without 0Auth
	r.POST("/register", func(c *gin.Context) {
		controller.Register(db, c)
	})
	// Register with 0Auth
	r.POST("/login", func(c *gin.Context) {
		controller.Login(db, c)
	})
	// Generate verification code
	r.POST("/generate", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.EmailVerificationCodeFromProfile(db, c)
	})
	// Verify email with verification code
	r.POST("verify", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.EmailVerificationNormal(db, c)
	})

	// Change password
	r.PATCH("/change-password", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.ChangePassword(db, c)
	})

	// Forgot password
	r.POST("/forgot-password", func(c *gin.Context) {
		controller.EmailVerificationCodeFromForgotPassword(db, c)
	})
	// Generate verification code for forgot password
	r.POST("/forgot-password/generate", func(c *gin.Context) {
		controller.EmailVerificationCodeFromForgotPassword(db, c)
	})
	// Verify verification code for forgot password
	r.POST("/forgot-password/verify", func(c *gin.Context) {
		controller.EmailVerificationForgotPassword(db, c)
	})

	// Login with 0Auth
	oauth := r.Group("/oauth")

	// Login with Google
	oauth.GET("/google", lib.GInit)
	oauth.GET("/google/callback", func(c *gin.Context) {
		controller.LoginOauth(db, c)
	})
}
