package server

import (
	"bytes"
	Database "kawan-usaha-api/db"
	"kawan-usaha-api/server/lib"
	"kawan-usaha-api/server/router"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	routerOnce sync.Once
	r          *gin.Engine
)

func SetupRouter() *gin.Engine {
	routerOnce.Do(func() {
		// DB
		db := Database.Open()
		if db != nil {
			log.Println("Nice, DB Connected")
		}

		// Gin Framework
		gin.SetMode(os.Getenv("GIN_MODE"))
		r = gin.Default()
		r.SetTrustedProxies(
			[]string{
				os.Getenv("PROXY_1"),
				os.Getenv("PROXY_2"),
				os.Getenv("PROXY_3"),
			},
		)

		//CORS
		r.Use(func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			if c.Request.Method == "OPTIONS" {
				c.Writer.Header().Set("Content-Type", "application/json")
				c.AbortWithStatus(204)
			} else {
				c.Next()
			}
		})

		// Config
		r.NoRoute(func(c *gin.Context) {
			c.JSON(404, lib.ErrorResponse("API Not Found", nil))
		})

		r.RemoveExtraSlash = true
		r.RedirectTrailingSlash = true
		r.Use(ginBodyLogMiddleware)

		//Routers

		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, lib.OkResponse("Welcome to Kawan Usaha API!", nil))
		})

		userInstance := router.GetUserInstance()
		userInstance.Init(db, r)
		userInstance.SetupRoutes()

		authInstance := router.GetAuthInstance()
		authInstance.Init(db, r)
		authInstance.SetupRoutes()

		usahaInstance := router.GetUsahaInstance()
		usahaInstance.Init(db, r)
		usahaInstance.SetupRoutes()

		articleInstance := router.GetArticleInstance()
		articleInstance.Init(db, r)
		articleInstance.SetupRoutes()

		categoryInstance := router.GetCategoryInstance()
		categoryInstance.Init(db, r)
		categoryInstance.SetupRoutes()

		chatInstance := router.GetChatInstance()
		chatInstance.Init(db, r)
		chatInstance.SetupRoutes()
	})
	return r
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()
	if statusCode >= 400 && gin.Mode() != "release" {
		log.Println("Response body: " + blw.body.String())
	}
}
