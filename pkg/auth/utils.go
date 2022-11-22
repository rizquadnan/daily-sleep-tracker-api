package auth

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func decryptString(input string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	
	if (err != nil ) {
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