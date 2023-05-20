package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GInit(c *gin.Context) {
	var GoogleAuth = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	url := GoogleAuth.AuthCodeURL("state", oauth2.AccessTypeOnline)
	c.Redirect(302, url)
}

func GCallback(c *gin.Context) string {
	var GoogleAuth = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	code := c.Query("code")
	tok, err := GoogleAuth.Exchange(context.Background(), code)
	if err != nil {
		log.Panic(err)
		c.JSON(500, ErrorResponse("Error while exchanging token", err.Error()))
	}
	client := GoogleAuth.Client(context.Background(), tok)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Panic(err)
		c.JSON(500, ErrorResponse("Error while getting user info", err.Error()))
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "")
	if err != nil {
		log.Panic("JSON parse error: ", err)
	}
	println(prettyJSON.String())
	return prettyJSON.String()
}
