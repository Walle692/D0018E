package user

import (
	"context"
	"fmt"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func CreateUser(cUsername string, cPassword string, cRole string) error {
	ctx := context.Background()

	postgres := global.Get()

	query := "INSERT INTO myschema.users (username, password, role) VALUES ($1, $2, $3)"

	_, err := postgres.Pool().Exec(ctx, query, cUsername, cPassword, cRole)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	// code to create a basket for the user, IF THEY ARE BUYER using the user id from the users table
	if cRole == "buyer" {
		// query to get the user id from the username
		var userID int
		query = "SELECT user_id FROM myschema.users WHERE username=$1"
		err = postgres.Pool().QueryRow(ctx, query, cUsername).Scan(&userID)
		if err != nil {
			return fmt.Errorf("create user get user id: %w", err)
		}

		// query to create a basket for the user
		query = "INSERT INTO myschema.basket (basket_user_id) VALUES ($1)"
		_, err = postgres.Pool().Exec(ctx, query, userID)
		if err != nil {
			return fmt.Errorf("create user create basket: %w", err)
		}
	}

	return nil

}
