package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// function to get the password from a specific username on the db
func (pg *Postgres) GetToken(ctx context.Context, username string) (string, error) {

	// variable to store the Stored password if the user exists
	var token string

	// query
	query := "SELECT token FROM myschema.token WHERE username=$1"

	// query the database and select the role for the username and store in role
	err := pg.db.QueryRow(ctx, query, username).Scan(&token)

	if err == pgx.ErrNoRows {
		// no user found
		fmt.Println("DEBUG: TOKEN ERR NO ROWS")
		return "", errors.New("No user found")

	} else if err != nil {
		// other error
		fmt.Println("DEBUG: TOKEN OTHER GET ERROR")
		return "", errors.New("Unhandled user error")
	}

	return token, nil

}