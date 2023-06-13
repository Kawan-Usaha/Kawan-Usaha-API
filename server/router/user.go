package router

// user.go

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	onceUser     sync.Once
	userInstance *UserSingleton
)

type UserSingleton struct {
	db *gorm.DB
	q  *gin.Engine
}

func GetUserInstance() *UserSingleton {
	onceUser.Do(func() {
		userInstance = &UserSingleton{}
	})
	return userInstance
}

func (u *UserSingleton) Init(db *gorm.DB, q *gin.Engine) {
	u.db = db
	u.q = q
}

func (u *UserSingleton) SetupRoutes() {
	r := u.q.Group("/user")
	// Get logged in User profile
	r.GET("/profile", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UserProfile(u.db, c)
	})
	// Update logged in User profile
	r.PATCH("/profile", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.UpdateUserProfile(u.db, c)
	})
	// Get favorite articles
	r.GET("/favorite-articles", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.GetFavoriteArticles(u.db, c)
	})
}
