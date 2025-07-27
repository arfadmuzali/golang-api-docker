package utils

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	secretKey := os.Getenv("AUTH_SECRET")
	claims := jwt.MapClaims{"username": username, "exp": time.Now().Add(time.Hour * 24).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
