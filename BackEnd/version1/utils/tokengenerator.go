package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supermegasecretkey") // Replace with a secure key

func GenerateJWT(username string) (string, error) {

	// key value pairs
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	// creating the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signing the token
	return token.SignedString(jwtKey)
}
