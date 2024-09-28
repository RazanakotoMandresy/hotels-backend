package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cz context.WithValue doesn't accept an type srting only custom type
		type user_uuid string
		var uuid user_uuid
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			respond(w, errorResponse{"no token String "}, http.StatusUnauthorized)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(JWT_SECRET), nil
		})

		if err != nil {
			respond(w, errorResponse{err.Error() + " error during parsing token"}, http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			respond(w, errorResponse{"token invalid "}, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			respond(w, errorResponse{" error's happen during claims"}, http.StatusUnauthorized)
			return
		}

		userUUID, ok := claims["user_uuid"].(string)
		if !ok {
			respond(w, errorResponse{" mo uuid during claims "}, http.StatusUnauthorized)
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
	// cz signed string already return err and string
	return token.SignedString([]byte(JWT_SECRET))
}
