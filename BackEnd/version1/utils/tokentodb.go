package utils

import (
	"context"
	"time"
)

// function to add the token along with expirytime to the db
func (pg *Postgres) TokenToDB(ctx context.Context, token string, expires_at time.Time, username string) error {

	// query
	query := "UPDATE myschema.test SET token=$1, token_expires_at=$2 WHERE uname=$3"

	// execute query
	_, err := pg.db.Exec(ctx, query, token, expires_at, username)
	if err != nil {
		return err
	}

	return nil
}

/* this is for use with pgx not pgx pool
func TokenToDB(token string, expires_at time.Time, username string, conn *pgx.Conn) error {

	query := "UPDATE myschema.test SET token=$1, token_expires_at=$2 WHERE uname=$3"

	_, err := conn.Exec(context.Background(), query, token, expires_at, username)
	if err != nil {
		return err
	}

	return nil
}
*/
