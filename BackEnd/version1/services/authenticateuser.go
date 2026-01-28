package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Walle692/D0018E/BackEnd/version1/utils"
)

func AuthenticateUser(pg *utils.Postgres, username, password string) (string, error) {

	// get password from the databse
	storedPassword, err := pg.GetPword(context.Background(), username)

	if password != storedPassword {
		// passwords didn't match
		fmt.Print("problema")
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, expires_at, err := utils.GenerateJWT(username)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	// add the generated token to the db passing along the db connection
	err = pg.TokenToDB(context.Background(), token, expires_at, username)
	if err != nil {
		return "", errors.New("failed to bind token")
	}

	return token, nil
}
