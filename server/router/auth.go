package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	onceAuth     sync.Once
	authInstance *AuthSingleton
)

type AuthSingleton struct {
	db *gorm.DB
	q  *gin.Engine
}

func GetAuthInstance() *AuthSingleton {
	onceAuth.Do(func() {
		authInstance = &AuthSingleton{}
	})
	return authInstance
}

func (a *AuthSingleton) Init(db *gorm.DB, q *gin.Engine) {
	a.db = db
	a.q = q
}

func (a *AuthSingleton) SetupRoutes() {
	r := a.q.Group("/auth")
	// Register without 0Auth
	r.POST("/register", func(c *gin.Context) {
		controller.Register(a.db, c)
	})
	// Register with 0Auth
	r.POST("/login", func(c *gin.Context) {
		controller.Login(a.db, c)
	})
	// Generate verification code
	r.POST("/generate", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.EmailVerificationCodeFromProfile(a.db, c)
	})
	// Verify email with verification code
	r.POST("/verify", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.EmailVerificationNormal(a.db, c)
	})
	// Change password
	r.POST("/change-password", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.ChangePassword(a.db, c)
	})
	// Forgot password
	r.POST("/forgot-password", func(c *gin.Context) {
		controller.EmailVerificationCodeFromForgotPassword(a.db, c)
	})

	s := r.Group("/forgot-password")
	// Generate verification code for forgot password
	s.POST("/generate", func(c *gin.Context) {
		controller.EmailVerificationCodeFromForgotPassword(a.db, c)
	})
	// Verify verification code for forgot password
	s.POST("/verify", func(c *gin.Context) {
		controller.EmailVerificationForgotPassword(a.db, c)
	})

	// Login with 0Auth
	oauth := r.Group("/oauth")
	// Login with Google
	oauth.GET("/google", lib.GInit)
	oauth.GET("/google/callback", func(c *gin.Context) {
		controller.LoginOauth(a.db, c)
	})
}
