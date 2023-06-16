package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/translate"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type ai_chat struct {
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

type ai_chat_continue struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Stream      bool    `json:"stream"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
}

type ai_chat_response_title struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index         int    `json:"index"`
		Finish_reason string `json:"finish_reason"`
		Message       struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		Prompt_tokens     int `json:"prompt_tokens"`
		Total_tokens      int `json:"total_tokens"`
		Completion_tokens int `json:"completion_tokens"`
	} `json:"usage"`
	Code int `json:"code"`
}

type ai_article_title struct {
	Article []string `json:"article"`
}

// error code is 40303

type ai_chat_response_article struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index         int    `json:"index"`
		Finish_reason string `json:"finish_reason"`
		Message       struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		Prompt_tokens     int `json:"prompt_tokens"`
		Total_tokens      int `json:"total_tokens"`
		Completion_tokens int `json:"completion_tokens"`
	} `json:"usage"`
	Code int `json:"code"`
}

func GetCompletions(db *gorm.DB, c *gin.Context) {
	var chat ai_chat

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
			c.JSON(http.StatusInternalServerError, lib.ErrorResponse("Failed to read the chat service response.", err.Error()))
			resp.Body.Close()
			return
		}

		// Send the complete data to the client
		c.Data(http.StatusOK, resp.Header.Get("Content-Type"), respData)
		resp.Body.Close()
	}
}

func ContinueCompletions(db *gorm.DB, c *gin.Context) {
	var chat ai_chat_continue

	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, lib.ErrorResponse("Failed to bind the request body.", err.Error()))
		return
	}

	jsonData, err := json.Marshal(chat)
	resp, err := http.Post(os.Getenv("LLM_URL")+"/v1/completions", "application/json", bytes.NewBuffer(jsonData))
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
			c.JSON(http.StatusInternalServerError, lib.ErrorResponse("Failed to read the chat service response.", err.Error()))
			resp.Body.Close()
			return
		}

		// Send the complete data to the client
		c.Data(http.StatusOK, resp.Header.Get("Content-Type"), respData)
		resp.Body.Close()
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
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse("Failed to read the chat service response.", err.Error()))
		resp.Body.Close()
		return
	}
	c.Data(http.StatusOK, resp.Header.Get("Content-Type"), respData)
	resp.Body.Close()
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
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse("Failed to read the chat service response.", err.Error()))
		resp.Body.Close()
		return
	}
	resp.Body.Close()
	c.Data(http.StatusOK, resp.Header.Get("Content-Type"), respData)
	resp.Body.Close()
}

func GenerateArticle(db *gorm.DB, c *gin.Context) {
	requestBody, err := c.GetRawData()
	if err != nil {
		log.Panicf("Failed to read the request body. %v", err.Error())
		return
	}

	go func() {
		sub, _ := c.Get("sub")
		subs := sub.(string)
		translator, _ := lib.GetTranslationClient()
		var chat ai_chat
		if err := json.Unmarshal(requestBody, &chat); err != nil {
			log.Printf("Failed to unmarshal the request body. %v", err.Error())
			return
		}
		var article_titles []string
		var summary string
		article_titles, summary = prepareArticlePrompt(db, c, translator, chat, subs)
		generateArticleContent(db, c, translator, chat, subs, article_titles, summary)
	}()

	c.JSON(http.StatusOK, lib.OkResponse("Generating article. Please wait for a few minutes", nil))
}

func prepareArticlePrompt(db *gorm.DB, c *gin.Context, translator *translate.Client, chat ai_chat, subs string) ([]string, string) {
	var article_titles []string
	var summary string
	// Create channels for communication
	articleTitlesChan := make(chan []string)
	summaryChan := make(chan string)
	go func() {
		i := 0
		for {
			var messages []struct {
				Role    string "json:\"role\""
				Content string "json:\"content\""
			}
			messages = chat.Messages[i:]
			jsonData, err := json.Marshal(messages)
			if err != nil {
				log.Printf("\n\n198 Failed to bind the request body. %v", err.Error())
				break
			}
			message := "Menggunakan kalimat berikut \"" + string(jsonData) + "\", Buatlah kesimpulan percakapan tersebut dalam satu kalimat."
			translated_message, _ := translator.Translate(context.Background(), []string{message}, language.AmericanEnglish, nil)

			new_chat := ai_chat{
				Model: chat.Model,
				Messages: []struct {
					Role    string "json:\"role\""
					Content string "json:\"content\""
				}{
					{
						Role:    "user",
						Content: translated_message[0].Text,
					},
				},
				Stream:      false,
				MaxTokens:   chat.MaxTokens,
				Temperature: chat.Temperature,
				TopP:        chat.TopP,
			}

			jsonData, err = json.Marshal(new_chat)
			if err != nil {
				log.Printf("\n\n225 Failed to bind the request body. %v", err.Error())
				break
			}

			resp, err := http.Post(os.Getenv("LLM_URL")+"/v1/chat/completions", "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				// Handle error
				log.Printf("\n\n232 Failed to communicate with the chat service. %v", err.Error())
				resp.Body.Close()
				break
			}
			defer resp.Body.Close()

			respData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// Handle error
				log.Printf("\n\n241 Failed to read the chat service response. %v", err.Error())
				resp.Body.Close()
				break
			}

			var respDataParsed ai_chat_response_title
			err = json.Unmarshal(respData, &respDataParsed)
			if err != nil {
				log.Printf("\n\n249 Failed to bind the request body. %v", err.Error())
				resp.Body.Close()
				break
			}

			if respDataParsed.Code > 0 {
				i++
				resp.Body.Close()
				continue
			}

			summary = respDataParsed.Choices[0].Message.Content
			break
		}
		message := "Utilizing this sentence \"" + summary + "\", make 3 article titles that's relevant for the user. do not explain anyting and strictly use this json format for the answer: {\"article\": [\"title_1\",\"title_2\", \"title_3\"]}"

		new_chat := ai_chat{
			Model: chat.Model,
			Messages: []struct {
				Role    string "json:\"role\""
				Content string "json:\"content\""
			}{
				{
					Role:    "user",
					Content: message,
				},
			},
			Stream:      false,
			MaxTokens:   chat.MaxTokens,
			Temperature: chat.Temperature,
			TopP:        chat.TopP,
		}

		jsonData, err := json.Marshal(new_chat)
		if err != nil {
			log.Printf("\n\n287 Failed to bind the request body. %v", err.Error())
			return
		}

		resp, err := http.Post(os.Getenv("LLM_URL")+"/v1/chat/completions", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			// Handle error
			log.Printf("\n\n294 Failed to communicate with the chat service. %v", err.Error())
			resp.Body.Close()
			return
		}
		defer resp.Body.Close()

		respData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// Handle error
			log.Printf("\n\n303 Failed to read the chat service response. %v", err.Error())
			resp.Body.Close()
			return
		}

		var respDataParsed ai_chat_response_title
		err = json.Unmarshal(respData, &respDataParsed)
		if err != nil {
			log.Printf("\n\n311 Failed to bind the request body. %v", err.Error())
			resp.Body.Close()
			return
		}

		var respDataParsedArticle ai_article_title
		err = json.Unmarshal([]byte(respDataParsed.Choices[0].Message.Content), &respDataParsedArticle)
		if err != nil {
			log.Printf("\n\n320 Failed to bind the request body. %v", err.Error())
			resp.Body.Close()
			return
		}

		if respDataParsed.Code > 0 {
			resp.Body.Close()
			return
		}

		if respDataParsedArticle.Article == nil {
			log.Printf("\n\n332 Failed to generate article title. %v", respDataParsed)
			resp.Body.Close()
			return
		} else {
			article_titles = respDataParsedArticle.Article
			articleTitlesChan <- article_titles
			summaryChan <- summary
			resp.Body.Close()
		}
	}()

	article_titles = <-articleTitlesChan
	summary = <-summaryChan

	return article_titles, summary
}

func generateArticleContent(db *gorm.DB, c *gin.Context, translator *translate.Client, chat ai_chat, subs string, article_titles []string, summary string) {
	go func() {
		if len(article_titles) == 0 {
			return
		}
		for _, article_title := range article_titles {
			new_chat := ai_chat{
				Model: chat.Model,
				Messages: []struct {
					Role    string "json:\"role\""
					Content string "json:\"content\""
				}{
					{
						Role:    "user",
						Content: "Utilizing this sentence \"" + summary + "\", make an article titles that's relevant for the user.",
					},
					{
						Role:    "assistant",
						Content: article_title,
					},
					{
						Role:    "user",
						Content: "write an article using that title. make sure to write an original article that is accurate and relevant to the title. write at least 1000 words.",
					},
				},
				Stream:      false,
				MaxTokens:   chat.MaxTokens,
				Temperature: chat.Temperature,
				TopP:        chat.TopP,
			}
			jsonData, err := json.Marshal(new_chat)
			if err != nil {
				log.Printf("\n\n368 Failed to bind the request body. %v", err.Error())
				return
			}

			resp, err := http.Post(os.Getenv("LLM_URL")+"/v1/chat/completions", "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				// Handle error
				log.Printf("\n\n375 Failed to communicate with the chat service. %v", err.Error())
				resp.Body.Close()
				return
			}
			defer resp.Body.Close()

			respData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// Handle error
				log.Printf("\n\n384 Failed to read the chat service response. %v", err.Error())
				resp.Body.Close()
				return
			}

			var respDataParsed ai_chat_response_article
			err = json.Unmarshal(respData, &respDataParsed)
			if err != nil {
				log.Printf("\n\n392 Failed to bind the request body. %v", err.Error())
				resp.Body.Close()
				return
			}

			if respDataParsed.Code > 0 {
				resp.Body.Close()
				return
			}

			judul_indonesianized, _ := translator.Translate(context.Background(), []string{article_title}, language.Indonesian, nil)
			artikel_indonesianized, _ := translator.Translate(context.Background(), []string{lib.FormatParagraph(respDataParsed.Choices[0].Message.Content)}, language.Indonesian, &translate.Options{Format: translate.Text})

			new_article := Model.Article{
				Title:       judul_indonesianized[0].Text,
				Content:     artikel_indonesianized[0].Text,
				IsPublished: false,
				UserID:      subs,
				Image:       "",
				CategoryID:  1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			if err := db.Create(&new_article).Error; err != nil {
				log.Printf("\n\n416 Failed to create article. %v", err.Error())
				resp.Body.Close()
				return
			}
			resp.Body.Close()
		}
		log.Println("GPU is free from generating article.")
	}()
}
