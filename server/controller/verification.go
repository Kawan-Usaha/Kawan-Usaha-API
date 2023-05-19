package controller

import (
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EmailVerificationCodeFromRegister(db *gorm.DB, c *gin.Context, user Model.User) {
	words := strings.Split(user.Name, " ")
	code := lib.GenerateEmailCode()
	emaildata := lib.EmailData{
		Subject:   "Email Verification",
		Code:      code,
		FirstName: words[0],
	}
	verification := Model.Verification{
		UserId:           user.UserId,
		VerificationCode: code,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	if err := db.Create(&verification).Error; err != nil {
		c.JSON(500, lib.ErrorResponse("Failed to create verification", err.Error()))
		return
	}
	lib.SendMailSingleReceiver(user.Email, &emaildata, lib.Templates["email_verification"])
	c.JSON(200, lib.OkResponse("Email verification sent", nil))
}

func EmailVerificationCodeFromProfile(db *gorm.DB, c *gin.Context) {
	sub, _ := c.Get("sub")
	subs := sub.(string)
	var user Model.User
	if err := db.Where("user_id = ?", subs).First(&user).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to get user", err.Error()))
		return
	}
	words := strings.Split(user.Name, " ")
	code := lib.GenerateEmailCode()
	emaildata := lib.EmailData{
		Subject:   "Email Verification",
		Code:      code,
		FirstName: words[0],
	}
	verification := Model.Verification{
		UserId:           user.UserId,
		VerificationCode: code,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	if err := db.Create(&verification).Error; err != nil {
		c.JSON(500, lib.ErrorResponse("Failed to create verification", err.Error()))
		return
	}
	lib.SendMailSingleReceiver(user.Email, &emaildata, lib.Templates["email_verification"])
	c.JSON(200, lib.OkResponse("Email verification sent", nil))
}

type forgotPassword struct {
	Email string `json:"email"`
}

func EmailVerificationCodeFromForgotPassword(db *gorm.DB, c *gin.Context) {
	var input forgotPassword
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Invalid input", err.Error()))
		return
	}
	var user Model.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Email not found", err.Error()))
		return
	}
	words := strings.Split(user.Name, " ")
	code := lib.GenerateEmailCode()
	emaildata := lib.EmailData{
		Subject:   "Email Verification",
		Code:      code,
		FirstName: words[0],
	}
	verification := Model.Verification{
		UserId:           user.UserId,
		VerificationCode: code,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	if err := db.Create(&verification).Error; err != nil {
		c.JSON(500, lib.ErrorResponse("Failed to create verification", err.Error()))
		return
	}
	lib.SendMailSingleReceiver(user.Email, &emaildata, lib.Templates["email_reset_password"])
	c.JSON(200, lib.OkResponse("Email verification sent", nil))
}

func EmailVerificationNormal(db *gorm.DB, c *gin.Context) {
	var input Model.Verification
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Invalid input", err.Error()))
		return
	}
	var verification Model.Verification
	if err := db.Where("verification_code = ?", input.VerificationCode).First(&verification).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Verification code not found", err.Error()))
		return
	}
	if verification.CreatedAt.Add(time.Minute * 10).Before(time.Now()) {
		c.JSON(400, lib.ErrorResponse("Verification code expired, please create a new One", nil))
		return
	}
	if err := db.Model(&Model.User{}).Where("user_id=?", verification.UserId).Update("Verified", true).Error; err != nil {
		c.JSON(500, lib.ErrorResponse("Failed to update verification", err.Error()))
		return
	}
	c.JSON(200, lib.OkResponse("Email verified", nil))
}

type resetPassword struct {
	VerificationCode string `json:"verification_code"`
	Password         string `json:"password"`
	PasswordConfirm  string `json:"password_confirm"`
}

func EmailVerificationForgotPassword(db *gorm.DB, c *gin.Context) {
	var input resetPassword
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Invalid input", err.Error()))
		return
	}
	var verification Model.Verification
	if err := db.Where("verification_code = ?", input.VerificationCode).First(&verification).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Verification code not found", err.Error()))
		return
	}
	if verification.CreatedAt.Add(time.Minute * 10).Before(time.Now()) {
		c.JSON(400, lib.ErrorResponse("Verification code expired, please create a new One", nil))
		return
	}
	var user Model.User
	if err := db.Where("user_id = ?", verification.UserId).First(&user).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("User not found", err.Error()))
		return
	}
	if input.Password != input.PasswordConfirm {
		c.JSON(400, lib.ErrorResponse("Password not match", nil))
		return
	}
	hashedpassword, _ := lib.HashPassword(input.Password)
	if err := db.Model(&Model.User{}).Where("user_id=?", verification.UserId).Update("Password", hashedpassword).Error; err != nil {
		c.JSON(500, lib.ErrorResponse("Failed to update verification", err.Error()))
		return
	}
	c.JSON(200, lib.OkResponse("Password changed successfully", nil))
}
