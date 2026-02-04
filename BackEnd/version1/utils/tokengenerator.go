package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY")) // Replace with a secure key

func GenerateJWT(username string, role string) (string, error) {

	// creating the basis for the token with username and role, to add other stuff
	// need to change the ClaimsCustom in that case
	// key value pairs
	claims := ClaimsCustom {
		username,
		role
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:	   "test",
			Subject:   username,
		}
	}

	// creating the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signing the token
	tokenstring, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}
