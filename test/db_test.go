package test

import (
	Database "kawan-usaha-api/db"
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDb(t *testing.T) {
	lib.EnvLoaderTest()
	db := Database.Open()
	assert.NotNil(t, db)

	// Test User
	newUser := Model.User{
		UserId:    "1234567890",
		Name:      "Hello",
		Username:  "hello",
		Email:     "hello@gmail.com",
		Password:  "12345678",
		Verified:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&newUser); err.Error != nil {
		log.Fatal(err.Error.Error())
	}
	searchUser := Model.User{}
	if err := db.Where("user_id = ?", "1234567890").First(&searchUser).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "1234567890", searchUser.UserId)

	// Test Usaha
	newUsaha := Model.Usaha{
		UsahaName: "HelloUsaha",
		Type:      "HelloType",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Model(&newUser).Association("Usaha").Append(&newUsaha); err != nil {
		log.Fatal(err.Error())
	}
	searchUsaha := Model.Usaha{}
	if err := db.Where("id = ?", 1).First(&searchUsaha).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, uint(1), searchUsaha.ID)

	// Test Article
	newArticle := Model.Article{
		UserId:      "1234567890",
		Title:       "HelloTitle",
		Content:     "HelloContent",
		Image:       "HelloImage",
		IsPublished: true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := db.Create(&newArticle); err.Error != nil {
		log.Fatal(err.Error.Error())
	}
	searchArticle := Model.Article{}
	if err := db.Where("title = ?", "HelloTitle").First(&searchArticle).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "HelloTitle", searchArticle.Title)

	// Test Category
	newCategory := Model.Category{
		Title:     "HelloTitle",
		Image:     "HelloImage",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&newCategory); err.Error != nil {
		log.Fatal(err.Error.Error())
	}
	searchCategory := Model.Category{}
	if err := db.Where("title = ?", "HelloTitle").First(&searchCategory).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "HelloTitle", searchCategory.Title)

	// Test Tag
	newTag := Model.Tag{
		Name: "HelloTag",
	}
	if err := db.Model(&newUsaha).Association("Tags").Append(&newTag); err != nil {
		log.Fatal(err.Error())
	}
	if err := db.Model(&newCategory).Association("Tags").Append(&newTag); err != nil {
		log.Fatal(err.Error())
	}
	searchTag := Model.Tag{}
	if err := db.Where("name = ?", "HelloTag").First(&searchTag).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "HelloTag", searchTag.Name)

	// Test Chat
	newChat := Model.Chat{
		ChatId:    "1234567890",
		UserId:    "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&newChat); err.Error != nil {
		log.Fatal(err.Error.Error())
	}
	searchChat := Model.Chat{}
	if err := db.Where("chat_id = ?", "1234567890").First(&searchChat).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "1234567890", searchChat.ChatId)

	// Test Message
	newMessage := Model.Message{
		ChatId:    "1234567890",
		UserId:    "1234567890",
		Message:   "HelloMessage",
		CreatedAt: time.Now(),
	}
	if err := db.Create(&newMessage); err.Error != nil {
		log.Fatal(err.Error.Error())
	}
	searchMessage := Model.Message{}
	if err := db.Where("message = ?", "HelloMessage").First(&searchMessage).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "HelloMessage", searchMessage.Message)

	// Test TempCode
	newTempCode := Model.TempCode{
		Email: "hello@gmail.com",
		Code:  "123456",
	}
	if err := db.Create(&newTempCode); err.Error != nil {
		log.Fatal(err.Error.Error())
	}
	searchTempCode := Model.TempCode{}
	if err := db.Where("email = ?", "hello@gmail.com").First(&searchTempCode).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "hello@gmail.com", searchTempCode.Email)

}
