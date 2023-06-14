package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"kawan-usaha-api/server/lib"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCompletions(db *gorm.DB, c *gin.Context) {
	var chat struct {
		Model    string `json:"model"`
		Messages []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
		Stream      bool    `json:"stream"`
		MaxTokens   int     `json:"max_tokens"`
		Temperature float64 `json:"temperature"`
		TopP        float64 `json:"top_p"`
	}

	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, lib.ErrorResponse("Failed to bind the request body.", err.Error()))
		return
	}

	jsonData, err := json.Marshal(chat)
	resp, err := http.Post(os.Getenv("LLM_URL")+"/v1/chat/completions", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse("Failed to communicate with the chat service.", err.Error()))
		return
	}
	defer resp.Body.Close()

	if chat.Stream {
		// If stream=true, stream the response data to the client
		c.Status(http.StatusOK)
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")

		// Stream the response data to the client
		for {
			// Read a chunk of data from the response body
			// and write it to the response writer
			buf := make([]byte, 4096)
			n, err := resp.Body.Read(buf)
			if err != nil {
				// Handle error or end of response
				c.Writer.Write(buf[:n])
				c.Writer.Flush()
				return
			}

			// Write the chunk of data to the response writer
			c.Writer.Write(buf[:n])
			c.Writer.Flush()
		}
	} else {
		// If stream=false, wait until the response is complete
		// and then send the complete data to the client
		respData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// Handle error
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read the chat service response.",
			})
			return
		}

		// Send the complete data to the client
		c.Data(http.StatusOK, resp.Header.Get("Content-Type"), respData)
	}
}

func TokenCheck(db *gorm.DB, c *gin.Context) {
	resp, err := http.Post(os.Getenv("LLM_URL")+"/v1/token_check", "application/json", c.Request.Body)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse("Failed to communicate with the chat service.", err.Error()))
		return
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read the chat service response.",
		})
		return
	}
	c.Data(http.StatusOK, resp.Header.Get("Content-Type"), respData)
}

func GetModels(db *gorm.DB, c *gin.Context) {
	resp, err := http.Get(os.Getenv("LLM_URL") + "/v1/models")
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse("Failed to communicate with the chat service.", err.Error()))
		return
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read the chat service response.",
		})
		return
	}
	c.Data(http.StatusOK, resp.Header.Get("Content-Type"), respData)
}
