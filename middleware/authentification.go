package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSectret = []byte(os.Getenv("JWT_SECRET"))

// return anle token string
func TokenManage(mail, uuid string) (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mail": mail,
		"uuid": uuid,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := tokens.SignedString(jwtSectret)
	if err != nil {
		return "", fmt.Errorf("%v's error on tokenManages😥", err)
	}
	return tokenString, nil
}
func RequireAuth(r *http.Request) error {
	tokenString, err := extractToken(r)
	if err != nil {
		return err
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return errors.New("token expired 🕰️")
		}
	}
	return nil
}
func extractToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", errors.New("no bearer token found 😔")
}
func ExtractTokenUUID(r *http.Request) (string, error) {
	tokenString, err := extractToken(r)
	if err != nil {
		return "", err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method 😔")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprint(claims["uuid"]), nil
	}
	return "", errors.New("an error occured during the uuid's extraction 😔")
}
