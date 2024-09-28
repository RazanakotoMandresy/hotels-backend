package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cz context.WithValue doesn't accept an type srting only custom type
		type user_uuid string
		var uuid user_uuid
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("secret"), nil
		})

		if err != nil {
			return
		}

		if !token.Valid {
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return
		}

		userUUID, ok := claims["user_uuid"].(string)
		if !ok {
			return
		}
		uuid = "user_uuid"
		ctx := context.WithValue(r.Context(), uuid, userUUID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CreateToken(userUUID, mail string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_uuid"] = userUUID
	claims["mail"] = mail
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("JWT_SECRET"))
}
