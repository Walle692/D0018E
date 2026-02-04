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
	if err != nil {
		return "", err
	}

	// check entered password with stored password
	if password != storedPassword {
		// passwords didn't match
		fmt.Print("problema")
		return "", errors.New("invalid password")
	}

	// get the role of the user
	role, err := pg.GetRole(context.Background, username)
	if err != nil {
		return "", err
	}

	// Generate JWT token using username and role
	token, err := utils.GenerateJWT(username, role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	// add the generated token to the db passing along the db connection
	err = pg.TokenToDB(context.Background(), token, username)
	if err != nil {
		return "", errors.New("failed to bind token")
	}

	return token, nil
}
