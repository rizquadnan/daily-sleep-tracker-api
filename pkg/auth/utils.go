package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func decryptString(input string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	
	if (err != nil ) {
		return "", err
	}

	return string(hashedPassword), err
}