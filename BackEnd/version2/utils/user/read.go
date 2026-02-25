package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

// function to get the password from a specific username on the db
func GetPassword(username string) (string, string, int, error) {
	ctx := context.Background()

	// variable to store the Stored password if the user exists
	var password string
	var role string
	var userID int

	// query
	query := "SELECT password, role, user_id FROM myschema.users WHERE username=$1"

	// get the db connection from global
	postgres := global.Get()

	// query the database and select the password from the username and store in password
	err := postgres.Pool().QueryRow(ctx, query, username).Scan(&password, &role, &userID)

	if err == pgx.ErrNoRows {
		fmt.Println("DEBUG: PASSWORD ERR NO ROWS")
		// no user found
		return "", "", 0, errors.New("No user found")
	} else if err != nil {
		fmt.Println("DEBUG: PASSWORD OTHER GET ERROR")
		fmt.Println(err)
		// other error
		return "", "", 0, errors.New("Unhandled user error")
	}

	return password, role, userID, nil

}
