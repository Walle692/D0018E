package services

import (
	"context"
	"errors"
	"time"

	"github.com/Walle692/D0018E/BackEnd/version1/utils"
)

// function to check if token exists in db and is still active
func (pg *Postgres) TokenChecker(ctx context.Context, token string, username string) (string, error) {

	// using the utils gettoken to get the stored token from db
	tokenfromdb, err := utils.GetToken(ctx, username)
	if err != nil {
		return "", err
	}

	


}
