package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


var jwtkey = []byte ("secret_galonku")

func GenerateToken(userId string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"exp":    jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // token berlaku selama 24 jam
	})
	tokenString, err := token.SignedString(jwtkey)
	return tokenString, err
}