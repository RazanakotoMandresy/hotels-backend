package middleware

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		fmt.Printf("Erreur lors du cryptage du mots de passe : %v \n", err)
		return ""
	}
	return string(bytes)
}
