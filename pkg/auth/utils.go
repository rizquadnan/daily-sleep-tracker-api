package auth

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func decryptString(input string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(password))
}

func generateToken(userId uint) (string, error) {
	tokenLifeSpan, err := strconv.Atoi(viper.Get("TOKEN_LIFE_SPAN").(string))
	apiSecret := viper.Get("API_SECRET").(string)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour + time.Duration(tokenLifeSpan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(apiSecret))
}

func IsTokenValid(c *gin.Context) bool {
	tokenString := extractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("adnankeren"), nil
	})

	if err != nil {
		return false
	} else {
		return true
	}
}

func extractToken(c *gin.Context) string {
	token := c.Query("token")

	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}
