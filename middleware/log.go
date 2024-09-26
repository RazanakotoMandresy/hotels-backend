package middleware

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	chans := make(chan []byte)
	go func() {
		bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
		chans <- bytes
	}()
	bytes := <-chans
	return string(bytes), nil
}
