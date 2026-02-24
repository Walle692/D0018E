package basket

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func CreateBasketItem(userID int, productID int, quantity int) error {
	ctx := context.Background()
	pool := global.Get().Pool()

	var dummy int
	insertInstead := true

	if err := pool.QueryRow(ctx, `
		SELECT 1
		FROM myschema.basketitem bi
		JOIN myschema.basket b ON b.basket_id = bi.basket_id
		WHERE b.basket_user_id = $1 AND bi.product_id =$2  
		`, userID, productID,
	).Scan(&dummy); err == pgx.ErrNoRows {
		insertInstead = false
	} else if err != nil {
		return err
	}

	if insertInstead {
		if _, err := pool.Exec(ctx, `
			UPDATE myschema.basketitem bi
			SET quantity = bi.quantity + $1
			FROM myschema.basket b
			WHERE b.basket_id = bi.basket_id 
				AND b.basket_user_id = $2 
				AND bi.product_id = $3
			`, quantity, userID, productID,
		); err != nil {
			return err
		}
		return nil

	}

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
