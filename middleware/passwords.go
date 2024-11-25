package middleware

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// encrypt encrypts plain text using the AES encryption algorithm and a key.

func Decrypt(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Printf("Erreur lors du cryptage du mots de passe : %v \n", err)
		return "", err
	}
	return string(bytes), nil
}
