package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// encrypt encrypts plain text using the AES encryption algorithm and a key.
func Encrypt(plainText string) (string, error) {
	key := os.Getenv("CRYPT_KEY")
	// Convert the key and plaintext to byte slices.
	keyBytes := []byte(key)
	plainTextBytes := []byte(plainText)

	// Create a new AES cipher block using the key.
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Create a new GCM (Galois/Counter Mode) cipher for the AES block.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a random nonce (number used once) for encryption.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the plaintext using the GCM cipher and the nonce.
	cipherText := gcm.Seal(nonce, nonce, plainTextBytes, nil)

	// Encode the cipherText to base64 to make it easier to handle.
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// decrypt decrypts the cipher text back to the original plain text using the AES encryption algorithm and a key.
func Decrypt(cipherText string) (string, error) {
	key := os.Getenv("CRYPT_KEY")
	// Convert the key and cipherText to byte slices.
	keyBytes := []byte(key)
	cipherTextBytes, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		fmt.Println(1)

		return "", err
	}

	// Create a new AES cipher block using the key.
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		fmt.Println(2)

		return "", err
	}

	// Create a new GCM cipher for the AES block.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(3)
		return "", err
	}

	// Extract the nonce from the cipherText.
	nonceSize := gcm.NonceSize()
	nonce, cipherTextBytes := cipherTextBytes[:nonceSize], cipherTextBytes[nonceSize:]

	// Decrypt the cipherText back to the original plaintext.
	plainTextBytes, err := gcm.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		fmt.Println(4)
		return "", err
	}

	return string(plainTextBytes), nil
}
