package tests

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/walle692/D0018E/BackEnd/version2/test_setup"
	"github.com/walle692/D0018E/BackEnd/version2/utils/basket"
)

/*
	Deleted the basket items and create them as Order items when all basket items are deleted
	Finnish building orderdate and total price
*/

func TestBasketToOrder(t *testing.T) {
	ctx, pool := test_setup.SetUpDB(t)

	sellerID := test_setup.SeedUser(t, ctx, pool, "seller_"+Suffix(), "seller")
	p1qty, p2qty := 100, 200

	t.Run("Out of stock case", func(t *testing.T) {
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		p1 := test_setup.SeedProduct(t, ctx, pool, sellerID, 100, p1qty)
		basketID := test_setup.SeedBasket(t, ctx, pool, buyerID)
		qty := p1qty + 1
		basketItemID := test_setup.SeedBasketItem(t, ctx, pool, basketID, p1, qty)

		orderID, err := basket.ConvertBasketToOrder(basketID)
		require.Error(t, err)
		require.Zero(t, orderID)
		//Check that order item wasn't deleted
		var count int
		pool.QueryRow(ctx, `SELECT COUNT(*) FROM myschema.basketitem WHERE basket_item_id=$1`, basketItemID).Scan(&count)
		require.NotZero(t, count)

		//Check that product stock is unchanged
		pool.QueryRow(ctx, `SELECT stock FROM myschema.products WHERE product_id=$1`, p1).Scan(&count)
		require.Equal(t, p1qty, count)
		t.Cleanup(func() { deleteNewOrder(t, ctx, pool, orderID) })

	})
	t.Run("1 basket item", func(t *testing.T) {
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		p1 := test_setup.SeedProduct(t, ctx, pool, sellerID, 100, p1qty)
		basketID := test_setup.SeedBasket(t, ctx, pool, buyerID)
		price, qty := 100.0, 10
		basketItemID := test_setup.SeedBasketItem(t, ctx, pool, basketID, p1, qty)

		orderID, err := basket.ConvertBasketToOrder(basketID)
		require.NoError(t, err)
		require.NotZero(t, orderID)
		//Test that basket item is removed from table
		var count int
		pool.QueryRow(ctx, `SELECT COUNT(*) FROM myschema.basketitem WHERE basket_item_id=$1`, basketItemID).Scan(&count)
		require.Equal(t, 0, count)
		//Test the order was created

		//Get the latest order_id from customer
		var orderUserID int
		var orderTotalPrice float64
		pool.QueryRow(ctx, `SELECT order_user_id, totalprice FROM myschema.order WHERE order_id=$1`, orderID).Scan(&orderUserID, &orderTotalPrice)
		require.Equal(t, buyerID, orderUserID)
		require.InDelta(t, price*float64(qty), orderTotalPrice, 0.1)

		//Check that it exists as order item
		pool.QueryRow(ctx, `SELECT COUNT(*) FROM myschema.orderitem WHERE order_id=$1`, orderID).Scan(&count)
		require.Equal(t, 1, count)

		//Check that product quantity is correctly reduced.
		pool.QueryRow(ctx, `SELECT stock FROM myschema.products WHERE product_id=$1`, p1).Scan(&count)
		require.Equal(t, p1qty-qty, count)

		t.Cleanup(func() { deleteNewOrder(t, ctx, pool, orderID) })

	})

	t.Run("2 basket item", func(t *testing.T) {
		p1 := test_setup.SeedProduct(t, ctx, pool, sellerID, 100, p1qty)
		p2 := test_setup.SeedProduct(t, ctx, pool, sellerID, 200, p2qty)
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		basketID := test_setup.SeedBasket(t, ctx, pool, buyerID)

		price_1, qty_1 := 100.0, 10
		basketItemID_1 := test_setup.SeedBasketItem(t, ctx, pool, basketID, p1, qty_1)

		price_2, qty_2 := 200.0, 20
		basketItemID_2 := test_setup.SeedBasketItem(t, ctx, pool, basketID, p2, qty_2)

		orderID, err := basket.ConvertBasketToOrder(basketID)
		require.NoError(t, err)
		require.NotZero(t, orderID)
		//Test that basket item is removed from table
		var count int
		pool.QueryRow(ctx, `SELECT COUNT(*) FROM myschema.basketitem WHERE basket_item_id=$1`, basketItemID_1).Scan(&count)
		require.Equal(t, 0, count)
		pool.QueryRow(ctx, `SELECT COUNT(*) FROM myschema.basketitem WHERE basket_item_id=$1`, basketItemID_2).Scan(&count)
		require.Equal(t, 0, count)
		//Test the order was created

		//Get the latest order_id from customer
		var orderUserID int
		var orderTotalPrice float64
		pool.QueryRow(ctx, `SELECT order_user_id, totalprice FROM myschema.order WHERE order_id=$1`, orderID).Scan(&orderUserID, &orderTotalPrice)
		require.Equal(t, buyerID, orderUserID)
		require.InDelta(t, price_1*float64(qty_1)+price_2*float64(qty_2), orderTotalPrice, 0.1)

		//Check that it exists as order item
		pool.QueryRow(ctx, `SELECT COUNT(*) FROM myschema.orderitem WHERE order_id=$1`, orderID).Scan(&count)
		require.Equal(t, 2, count)

		//Check that product quantity is correctly reduced.
		pool.QueryRow(ctx, `SELECT stock FROM myschema.products WHERE product_id=$1`, p1).Scan(&count)
		require.Equal(t, p1qty-qty_1, count)
		pool.QueryRow(ctx, `SELECT stock FROM myschema.products WHERE product_id=$1`, p2).Scan(&count)
		require.Equal(t, p2qty-qty_2, count)

		t.Cleanup(func() { deleteNewOrder(t, ctx, pool, orderID) })
	})

	t.Run("Product no longer active", func(t *testing.T) {
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		p1 := test_setup.SeedProduct(t, ctx, pool, sellerID, 100, p1qty)
		basketID := test_setup.SeedBasket(t, ctx, pool, buyerID)
		qty := 10
		_ = test_setup.SeedBasketItem(t, ctx, pool, basketID, p1, qty)

		pool.Exec(ctx, `UPDATE myschema.products SET active = FALSE WHERE product_id=$1`, p1)

		orderID, err := basket.ConvertBasketToOrder(basketID)
		require.Error(t, err)
		t.Cleanup(func() { deleteNewOrder(t, ctx, pool, orderID) })

	})

}

func deleteNewOrder(t *testing.T, ctx context.Context, pool *pgxpool.Pool, orderID int) {
	if _, err := pool.Exec(ctx, `DELETE FROM myschema.orderitem WHERE order_id=$1`, orderID); err != nil {
		t.Logf("clean up order item: %v", err)
	}
	if _, err := pool.Exec(ctx, `DELETE FROM myschema.order WHERE order_id=$1`, orderID); err != nil {

		t.Logf("clean up order: %v", err)
	}
}
