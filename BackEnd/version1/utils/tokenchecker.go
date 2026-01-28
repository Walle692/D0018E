package utils

import (
	"context"
	"errors"
	"time"
)

// function to add the token along with expirytime to the db
func (pg *Postgres) TokenChecker(ctx context.Context, token string) (string, error) {

	// query
	query := "SELECT (uname, token_expires_at) FROM myschema.test WHERE token=$1"

	var username string
	var expires_at time.Time

	// query the database and select the username and expiry data
	err := pg.db.QueryRow(ctx, query, token).Scan(&username, &expires_at)
	if err != nil {
		return "", err
	}

	// check that the token has not expired
	if time.Now().After(expires_at) {
		return "", errors.New("Token expired")
	}

	return username, nil
}
