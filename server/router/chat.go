package router

import (
	"kawan-usaha-api/server/controller"
	"kawan-usaha-api/server/lib"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	onceChat     sync.Once
	chatInstance *ChatSingleton
)

type ChatSingleton struct {
	db *gorm.DB
	q  *gin.Engine
}

func GetChatInstance() *ChatSingleton {
	onceChat.Do(func() {
		chatInstance = &ChatSingleton{}
	})
	return chatInstance
}

func (c *ChatSingleton) Init(db *gorm.DB, q *gin.Engine) {
	c.db = db
	c.q = q
}

func (chat *ChatSingleton) SetupRoutes() {
	r := chat.q.Group("/v1")
	r.POST("/chat/completions", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.GetCompletions(chat.db, c)
	})
	r.POST("/token_check", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.TokenCheck(chat.db, c)
	})
	r.GET("/models", lib.ValidateJWTToken(), func(c *gin.Context) {
		controller.GetModels(chat.db, c)
	})
}
