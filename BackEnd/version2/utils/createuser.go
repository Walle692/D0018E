package utils

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
	return nil

}
