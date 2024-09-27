package middleware

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	chans := make(chan []byte)
	go func() {
		bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
		chans <- bytes
	}()
	bytes := <-chans
	return string(bytes)
}
func VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
