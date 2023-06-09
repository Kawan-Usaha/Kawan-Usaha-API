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
		UserId:    "59d729af-f5a6-4e6c-9eac-027ed3fc11e0",
		Name:      "Hello",
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
	if err := db.Where("user_id = ?", "59d729af-f5a6-4e6c-9eac-027ed3fc11e0").First(&searchUser).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "59d729af-f5a6-4e6c-9eac-027ed3fc11e0", searchUser.UserId)

	// Test Usaha
	newUsaha := Model.Usaha{
		UsahaName: "HelloUsaha",
		Type:      1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Model(&newUser).Association("Usaha").Append(&newUsaha); err != nil {
		log.Fatal(err.Error())
	}
	searchUsaha := Model.Usaha{}
	if err := db.Where("usaha_name = ?", "HelloUsaha").First(&searchUsaha).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "HelloUsaha", searchUsaha.UsahaName)

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

	// Test Article
	newArticle := Model.Article{
		UserID:      "1234567890",
		Title:       "HelloTitle",
		Content:     "HelloContent",
		Image:       "HelloImage",
		IsPublished: true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	newArticle.Category = newCategory
	newArticle.User = newUser
	if err := db.Create(&newArticle).Error; err != nil {
		log.Fatal(err.Error())
	}

	searchArticle := Model.Article{}
	if err := db.Where("title = ?", "HelloTitle").First(&searchArticle).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "HelloTitle", searchArticle.Title)

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

	// Test Verification
	newVerification := Model.Verification{
		VerificationCode: "123456",
	}
	if err := db.Model(&newUser).Association("Verification").Append(&newVerification); err != nil {
		log.Fatal(err.Error())
	}
	searchVerification := Model.Verification{}
	if err := db.Where("user_id = ?", "59d729af-f5a6-4e6c-9eac-027ed3fc11e0").First(&searchVerification).Error; err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, "59d729af-f5a6-4e6c-9eac-027ed3fc11e0", searchVerification.UserID)
	//remove existing data
	db.Model(&Model.User{}).Association("Usaha").Clear()
	db.Model(&Model.User{}).Association("Article").Clear()
	db.Model(&Model.Usaha{}).Association("Tags").Clear()
	db.Model(&Model.Category{}).Association("Tags").Clear()
	db.Model(&Model.Article{}).Association("Category").Clear()
	db.Model(&Model.Article{}).Association("User").Clear()
	db.Model(&Model.Tag{}).Association("Usaha").Clear()
	db.Model(&Model.Tag{}).Association("Category").Clear()

	db.Delete(&Model.Tag{}, "id = ?", uint(1))
	db.Delete(&Model.Verification{}, "user_id = ?", "59d729af-f5a6-4e6c-9eac-027ed3fc11e0")
	db.Delete(&Model.Article{}, "id = ?", uint(1))
	db.Delete(&Model.Category{}, "id = ?", uint(1))
	db.Delete(&Model.Usaha{}, "id = ?", uint(1))
	db.Delete(&Model.User{}, "user_id = ?", "59d729af-f5a6-4e6c-9eac-027ed3fc11e0")
}
