package middleware

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return "", fmt.Errorf("an error during the encryption of the passwords : %v ", err)

	}
	return string(bytes), err
}
