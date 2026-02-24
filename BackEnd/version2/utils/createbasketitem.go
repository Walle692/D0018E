package utils

import (
	"context"
	"fmt"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func CreateBasketItem(userID int, productID int, quantity int) error {
	ctx := context.Background()

	pool := global.Get().Pool()

	// insert the basketitem into the db
	query := `
	INSERT INTO myschema.basketitem (basket_id, product_id, quantity)
	SELECT b.basket_id, $2, $3
	FROM myschema.basket b
	WHERE b.basket_user_id = $1
	`
	_, err := pool.Exec(ctx, query, userID, productID, quantity)
	if err != nil {
		fmt.Println("DEBUG: CREATE BASKET ITEM INSERT ERROR")
		return err
	}

	return nil
}
