package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// function to get the role for user from db
func (pg *Postgres) GetRole(ctx context.Context, username string) (string, error) {

	// variable to store the role
	var role string

	// query
	query := "SELECT role FROM myschema.users WHERE username=$1"

	// query the database and select the role for the username and store in role
	err := pg.db.QueryRow(ctx, query, username).Scan(&role)

	if err == pgx.ErrNoRows {
		// no user found
		fmt.Println("DEBUG: ROLE ERR NO ROWS")
		return "", errors.New("No user found")

	} else if err != nil {
		// other error
		fmt.Println("DEBUG: ROLE OTHER GET ERROR")
		return "", errors.New("Unhandled user error")
	}

	return role, nil

}
