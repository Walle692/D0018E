package utils

import (
	"context"
	"time"
)

// function to add the token along with expirytime to the db
func (pg *Postgres) TokenToDB(ctx context.Context, token string, expires_at time.Time, username string) error {

	// query
	query := "UPDATE myschema.token SET token=$1 WHERE username=$3"

	// execute query
	_, err := pg.db.Exec(ctx, query, token, expires_at, username)
	if err != nil {
		return err
	}

	return nil
}
