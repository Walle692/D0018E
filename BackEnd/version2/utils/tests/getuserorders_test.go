package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/walle692/D0018E/BackEnd/version2/test_setup"
	"github.com/walle692/D0018E/BackEnd/version2/utils/order"
)

/*
Test that function correctly returns none, one, and two orders
*/
func TestRootOrders(t *testing.T) {
	ctx, pool := test_setup.SetUpDB(t)

	sellerID := test_setup.SeedUser(t, ctx, pool, "seller_"+Suffix(), "seller")
	productID := test_setup.SeedProduct(t, ctx, pool, sellerID, 100, 10)

	t.Run("no orders", func(t *testing.T) {
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		out, err := order.GetUserOrders(buyerID)
		require.NoError(t, err)
		require.Empty(t, out)
	})

	t.Run("one order", func(t *testing.T) {
		totalPrice := 100.0
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		orderID := test_setup.SeedOrder(t, ctx, pool, buyerID, totalPrice)
		_ = test_setup.SeedOrderItem(t, ctx, pool, orderID, productID, 1, 1)

		out, err := order.GetUserOrders(buyerID)
		require.NoError(t, err)
		require.Len(t, out, 1)

		require.Equal(t, orderID, out[0].Order_id)
		require.Equal(t, totalPrice, out[0].Total_price)
	})

	t.Run("two orders", func(t *testing.T) {
		totalPrice_1, totalPrice_2 := 100.0, 200.0
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		order1ID := test_setup.SeedOrder(t, ctx, pool, buyerID, totalPrice_1)
		_ = test_setup.SeedOrderItem(t, ctx, pool, order1ID, productID, 1, 1)
		order2ID := test_setup.SeedOrder(t, ctx, pool, buyerID, totalPrice_2)
		_ = test_setup.SeedOrderItem(t, ctx, pool, order2ID, productID, 1, 1)

		out, err := order.GetUserOrders(buyerID)
		require.NoError(t, err)
		require.Len(t, out, 2)

		byID := map[int]order.OrderResponse{}
		for _, o := range out {
			byID[o.Order_id] = o
		}

		require.Equal(t, totalPrice_1, byID[order1ID].Total_price)
		require.Equal(t, totalPrice_2, byID[order2ID].Total_price)
	})
}

/*
Test that a order correctly contains product and quantity, checks 1 and 2 products
*/
func TestOrderItems(t *testing.T) {
	ctx, pool := test_setup.SetUpDB(t)

	sellerID := test_setup.SeedUser(t, ctx, pool, "seller_"+Suffix(), "seller")
	// create 2 products you can reuse
	p1 := test_setup.SeedProduct(t, ctx, pool, sellerID, 100, 10)
	p2 := test_setup.SeedProduct(t, ctx, pool, sellerID, 200, 10)

	t.Run("1 order item", func(t *testing.T) {
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		orderID := test_setup.SeedOrder(t, ctx, pool, buyerID, 1)
		_ = test_setup.SeedOrderItem(t, ctx, pool, orderID, p1, 100, 100)

		out, err := order.GetUserOrders(buyerID)
		require.NoError(t, err)

		require.Len(t, out[0].Order_items, 1)
	})
	t.Run("2 order item", func(t *testing.T) {
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		orderID := test_setup.SeedOrder(t, ctx, pool, buyerID, 1)
		_ = test_setup.SeedOrderItem(t, ctx, pool, orderID, p1, 100, 100)
		_ = test_setup.SeedOrderItem(t, ctx, pool, orderID, p2, 200, 200)

		out, err := order.GetUserOrders(buyerID)
		require.NoError(t, err)

		require.Len(t, out[0].Order_items, 2)

		items := out[0].Order_items
		require.Equal(t, items[0].Quantity, 100)
		require.Equal(t, items[1].Quantity, 200)

		require.Equal(t, items[0].Product.Product_id, p1)
		require.Equal(t, items[1].Product.Product_id, p2)

	})
}

func Suffix() string {
	return fmt.Sprint(time.Now().UnixNano())
}
