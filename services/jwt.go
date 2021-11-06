package services

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToken(userID interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Authorized": true,
		"expiresAt": time.Now().Add(time.Hour * 1).Unix(),
		"issuedAt": time.Now().Unix(),
		"userId": userID,
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyToken(token string) (*jwt.Token, error){
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
