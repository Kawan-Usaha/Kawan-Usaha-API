package controller

import (
	Model "kawan-usaha-api/model"
	"kawan-usaha-api/server/lib"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, c *gin.Context) {
	var input Model.Register
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Invalid input", err.Error()))
		return
	}
	hashedpassword, err := lib.HashPassword(input.Password)
	if err != nil {
		c.JSON(500, lib.ErrorResponse("Failed to hash password", err.Error()))
		return
	}
	if input.Password != input.PasswordConfirm {
		c.JSON(403, lib.ErrorResponse("Password and password confirm doesn't match", nil))
		return
	}
	regist := Model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedpassword,
		UserId:   uuid.New().String(),
	}
	if err := db.Create(&regist).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Failed to register", err.Error()))
		return
	}
	EmailVerificationCodeFromRegister(db, c, regist)

	hours, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token, err := lib.GenerateJWTToken(time.Duration(hours)*time.Hour, regist.UserId)
	if err != nil {
		c.JSON(500, lib.ErrorResponse("Failed to generate token", err.Error()))
		return
	}

	result := gin.H{
		"token": token,
		"name":  regist.Name,
		"email": regist.Email,
	}

	c.JSON(200, lib.OkResponse("Register success, please login then validate your email address", result))
}

func Login(db *gorm.DB, c *gin.Context) {
	var input Model.Login
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, lib.ErrorResponse("Invalid input", err.Error()))
		return
	}
	var user Model.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(400, lib.ErrorResponse("Email or password is wrong", err.Error()))
		return
	}
	if err := lib.VerifyPassword(user.Password, input.Password); err != nil {
		c.JSON(400, lib.ErrorResponse("Email or password is wrong", err.Error()))
		return
	}
	hours, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token, err := lib.GenerateJWTToken(time.Duration(hours)*time.Hour, user.UserId)
	if err != nil {
		c.JSON(500, lib.ErrorResponse("Failed to generate token", err.Error()))
		return
	}
	result := gin.H{
		"token": token,
		"name":  user.Name,
		"email": user.Email,
	}
	c.JSON(200, lib.OkResponse("Login success", result))
}
