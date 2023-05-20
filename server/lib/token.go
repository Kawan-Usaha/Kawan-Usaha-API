package lib

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(ttl time.Duration, payload interface{}) (string, error) {
	// token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
	// 	"id":  user.UserID,
	// 	"exp": time.Now().Add(ttl).Unix(),
	// 	"iat": time.Now().Unix(),
	// 	"nbf": time.Now().Unix(),
	// })

	// tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	token := jwt.New(jwt.SigningMethodHS512)

	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}

func GenerateEmailCode() string {
	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

func ValidateJWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = header[len("Bearer "):]
		token, err := jwt.Parse(header, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})
		if err != nil {
			c.JSON(500, ErrorResponse("JWT invalid.", err.Error()))
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("sub", claims["sub"])
			c.Next()
			return
		} else {
			c.JSON(403, ErrorResponse("JWT invalid.", err.Error()))
			c.Abort()
			return
		}
	}
}
