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

/*
func AuthenticateUser(username, password string) (string, error) {
	// set up connection to database, the DATABASE url should be stored in env
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	// if error
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database %v\n", err)
		os.Exit(1)
	}

	// makes sure that connection is closed after authenticate user is done running (closes on return statement)
	defer conn.Close(context.Background())

	// variable to store the Stored password if the user exists
	var storedPassword string

	// query the database (mby use query instead of queryrow?) and select the password from the username and store in storedpassword
	err = conn.QueryRow(context.Background(), "SELECT pword FROM myschema.test WHERE uname=$1", username).Scan(&storedPassword)

	if err == pgx.ErrNoRows {
		// no user found
		return "", errors.New("No user found")
	} else if err != nil {
		// other error
	}

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
	err = TokenToDB(token, expires_at, username, conn)
	if err != nil {
		return "", errors.New("failed to bind token")
	}

	return token, nil
}
*/
