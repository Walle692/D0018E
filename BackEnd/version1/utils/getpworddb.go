package utils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

// function to get the password from a specific username on the db
func (pg *Postgres) GetPword(ctx context.Context, username string) (string, error) {

	// variable to store the Stored password if the user exists
	var storedPassword string

	// query
	query := "SELECT pword FROM myschema.test WHERE uname=$1"

	// query the database and select the password from the username and store in storedpassword
	err := pg.db.QueryRow(ctx, query, username).Scan(&storedPassword)

	if err == pgx.ErrNoRows {
		// no user found
		return "", errors.New("No user found")
	} else if err != nil {
		// other error
	}

	return storedPassword, nil

}
