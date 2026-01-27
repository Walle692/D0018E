package services

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func TokenToDB(token string, expires_at time.Time, username string, conn *pgx.Conn) error {

	query := "UPDATE myschema.test SET token=$1, token_expires_at=$2 WHERE uname=$3"

	_, err := conn.Exec(context.Background(), query, token, expires_at, username)
	if err != nil {
		return err
	}

	return nil
}
