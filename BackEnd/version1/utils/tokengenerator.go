package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY")) // Replace with a secure key

func GenerateJWT(username string) (string, time.Time, error) {
	expiry := time.Now().Add(time.Hour * 1)

	// key value pairs
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expiry.Unix(),
	}

	// creating the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signing the token
	tokenstring, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenstring, expiry, nil
}
