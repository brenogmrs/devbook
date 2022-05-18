package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 24).Unix()
	permissions["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte("secret"))
}