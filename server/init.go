package server

import (
	Database "kawan-usaha-api/db"
	"kawan-usaha-api/server/lib"
	"kawan-usaha-api/server/router"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// DB
	db := Database.Open()
	if db != nil {
		println("Nice, DB Connected")
	}

	// Gin Framework
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
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

	//Routers

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, lib.OkResponse("Welcome to Kawan Usaha API!", nil))
	})
	router.User(db, r)
	router.Auth(db, r)
	return r
}
