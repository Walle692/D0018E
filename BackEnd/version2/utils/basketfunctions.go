package utils

import (
	"context"
	"fmt"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

/*
Currently takes basketID for consistency among how other basket function are written can easily be
rewritten to take buyerID so we don't have to reverse the search done in getUserBasketID
*/
func ConvertBasketToOrder(basketID int) (orderID int, err error) {
	ctx := context.Background()
	pool := global.Get().Pool()

	tx, err := pool.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	// Get buyer ID (replace with get basket id if rewrite function input to buyerID)
	var buyerID int
	if err = tx.QueryRow(ctx, `SELECT basket_user_id FROM myschema.basket WHERE basket_id=$1`, basketID).Scan(&buyerID); err != nil {
		return 0, fmt.Errorf("basket not found: %w", err)
	}

	// Prepare total price from basket items
	var total float64
	var hasInActive bool
	if err = tx.QueryRow(ctx,
		`
		SELECT COALESCE(SUM(bi.quantity * p.price), 0), BOOL_OR(p.active = false)
		FROM myschema.basketitem bi
		JOIN myschema.products p ON p.product_id = bi.product_id
		WHERE basket_id=$1
		`, basketID,
	).Scan(&total, &hasInActive); err != nil {
		return 0, fmt.Errorf("calc total: %w", err)
	}

	if total == 0 {
		return 0, fmt.Errorf("basket is empty")
	}
	if hasInActive {
		return 0, fmt.Errorf("product not available")
	}

	// Perhaps return all products that doesn't have sufficient stock to better handle
	// there error when stock is not enough. But for now just abort and return error
	var errNoStock int
	if err = tx.QueryRow(ctx, `
		WITH cte AS (
			UPDATE myschema.products p
			SET stock = stock - bi.quantity
			FROM myschema.basketitem bi
			WHERE bi.basket_id = $1 AND p.product_id = bi.product_id
			RETURNING p.product_id, p.stock
		)
		SELECT product_id FROM cte WHERE stock < 0
		LIMIT 1
		`, basketID,
	).Scan(&errNoStock); err == nil {
		return 0, fmt.Errorf("insufficient stock product id %d", errNoStock)
	}
	err = nil

	// Build order
	if err = tx.QueryRow(ctx,
		`
		INSERT INTO myschema.order (order_user_id, orderdate, totalprice)
		VALUES ($1, NOW(), $2)
		RETURNING order_id
		`,
		buyerID, total,
	).Scan(&orderID); err != nil {
		return 0, fmt.Errorf("insert order: %w", err)
	}

	// Convert basket items to order items
	_, err = tx.Exec(ctx,
		`
		INSERT INTO myschema.orderitem (order_id, product_id, quantity, price)
		SELECT $1, bi.product_id, bi.quantity, p.price
		FROM myschema.basketitem bi
		JOIN myschema.products p ON p.product_id = bi.product_id 
		WHERE basket_id=$2
		`,
		orderID, basketID,
	)
	if err != nil {
		return 0, fmt.Errorf("insert order items: %w", err)
	}
	_, err = tx.Exec(ctx, `DELETE FROM myschema.basketitem WHERE basket_id=$1`, basketID)
	if err != nil {
		return 0, fmt.Errorf("delete basket items: %w", err)
	}

	return orderID, nil
}

func DeleteBasketItemsWithProduct(userID int, productID int) error {
	ctx := context.Background()
	pool := global.Get().Pool()

	if _, err := pool.Exec(ctx, `
		DELETE FROM myschema.basketitem bi
		USING myschema.basket b
		WHERE b.basket_id = bi.basket_id AND product_id = $1 AND basket_user_id = $2

		`, productID, userID,
	); err != nil {
		return err
	}
	return nil

}

//TODO DELETE BASKET ITEMS
