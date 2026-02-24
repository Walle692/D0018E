package basket

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

// function to get the basktet id from a specific userID on the db
func GetBasketID(c *gin.Context) (int, error) {
	ctx := context.Background()
	var basketID int

	// get the session from the request context
	session := sessions.Default(c)

	// get the username
	if user := session.Get(global.Userkey); user == nil {
		// no user in session, abort the request
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return 0, errors.New("Unauthorized")
	}

	// get the username string
	username := session.Get(global.Userkey).(string)

	// query to get the basket id using the username, joining the basket and userstables.
	query := "SELECT basket_id FROM myschema.basket LEFT JOIN myschema.users ON myschema.basket.basket_user_id=myschema.users.user_id WHERE username=$1"

	// get the db connection from global
	postgres := global.Get()

	// query the database and select the basket id from the userid
	err := postgres.Pool().QueryRow(ctx, query, username).Scan(&basketID)

	if err == pgx.ErrNoRows {
		fmt.Println("DEBUG: BASKETID ERR NO ROWS")
		// no user found
		return 0, errors.New("No basket found")
	} else if err != nil {
		fmt.Println("DEBUG: BAKETID OTHER GET ERROR")
		fmt.Println(err)
		// other error
		return 0, errors.New("Unexpected error")
	}

	return basketID, nil

}
