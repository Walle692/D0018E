package test_setup

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func SeedUser(t *testing.T, ctx context.Context, pool *pgxpool.Pool, username string, role string) (userID int) {
	t.Helper()

	userQuery := `--sql
		INSERT INTO myschema.users (username, password, role) 
		VALUES ($1, $2, $3)
		RETURNING user_id
		`
	err := pool.QueryRow(ctx, userQuery, username, username, role).Scan(&userID)
	require.NoError(t, err)

	t.Cleanup(func() {
		if _, e := pool.Exec(ctx, `DELETE FROM myschema.users WHERE user_id=$1`, userID); e != nil {
			t.Logf("clean up %s user: %v", username, e)
		}
	})

	return userID
}

func SeedProduct(t *testing.T, ctx context.Context, pool *pgxpool.Pool, sellerID int, price float64, stock int) (productID int) {
	t.Helper()

	n := time.Now().UnixNano()
	suffix := strconv.FormatInt(n, 10)

	Name, Manufacturer, Description, Screen_Size, Picture_url, Sku :=
		"Name_"+suffix, "Manufacturer_"+suffix, "Description_"+suffix, int(n), "Picture_url_"+suffix, "Sku_"+suffix

	prodQuery := `--sql
		INSERT INTO myschema.products (product_name, manufacturer ,seller_user_id, description, screen_size ,picture_url, sku, price, stock)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING product_id
	`
	err := pool.QueryRow(ctx, prodQuery, Name, Manufacturer, sellerID, Description, Screen_Size, Picture_url, Sku, price, stock).Scan(&productID)
	require.NoError(t, err)

	t.Cleanup(func() {
		if _, e := pool.Exec(ctx, `DELETE FROM myschema.products WHERE product_id=$1`, productID); e != nil {
			t.Logf("clean up product with %d : %v", sellerID, e)
		}
	})

	return productID
}

func SeedOrder(t *testing.T, ctx context.Context, pool *pgxpool.Pool, buyerID int, totalPrice float64) (orderID int) {
	t.Helper()

	orderQuery := `--sql
		INSERT INTO myschema.order (order_user_id, orderdate, totalprice)
		VALUES ($1, NOW(), $2)
		RETURNING order_id
	`
	err := pool.QueryRow(ctx, orderQuery, buyerID, totalPrice).Scan(&orderID)
	require.NoError(t, err)

	t.Cleanup(func() {
		if _, e := pool.Exec(ctx, `DELETE FROM myschema.order WHERE order_id=$1`, orderID); e != nil {
			t.Logf("clean up order: %v", e)
		}
	})
	return orderID
}

func SeedOrderItem(t *testing.T, ctx context.Context, pool *pgxpool.Pool, orderID int, productID int, quantity int, price float64) (orderItemID int) {
	t.Helper()

	orderItemQuery := `--sql
		INSERT INTO myschema.orderitem (order_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4)
		RETURNING order_item_id
	`
	err := pool.QueryRow(ctx, orderItemQuery, orderID, productID, quantity, price).Scan(&orderItemID)
	require.NoError(t, err)

	t.Cleanup(func() {
		if _, e := pool.Exec(ctx, `DELETE FROM myschema.orderitem WHERE order_item_id=$1`, orderItemID); e != nil {
			t.Logf("clean up order item: %v", e)
		}
	})

	return orderItemID
}

func SeedBasket(t *testing.T, ctx context.Context, pool *pgxpool.Pool, buyerID int) (basketID int) {
	t.Helper()

	basketQuery := `--sql
		INSERT INTO myschema.basket (basket_user_id)
		VALUES ($1)
		RETURNING basket_id
	`
	err := pool.QueryRow(ctx, basketQuery, buyerID).Scan(&basketID)
	require.NoError(t, err)

	t.Cleanup(func() {
		if _, e := pool.Exec(ctx, `DELETE FROM myschema.basket WHERE basket_id=$1`, basketID); e != nil {
			t.Logf("clean up basket: %v", e)
		}
	})
	return basketID
}

func SeedBasketItem(t *testing.T, ctx context.Context, pool *pgxpool.Pool, basketID int, productID int, quantity int, price float64) (basketItemID int) {
	t.Helper()

	basketItemQuery := `--sql
		INSERT INTO myschema.basketitem (basket_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4)
		RETURNING basket_item_id
	`
	err := pool.QueryRow(ctx, basketItemQuery, basketID, productID, quantity, price).Scan(&basketItemID)
	require.NoError(t, err)

	t.Cleanup(func() {
		if _, e := pool.Exec(ctx, `DELETE FROM myschema.basketitem WHERE basket_item_id=$1`, basketItemID); e != nil {
			t.Logf("clean up basket item: %v", e)
		}
	})

	return basketItemID
}
