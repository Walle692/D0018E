package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/walle692/D0018E/BackEnd/version2/test_setup"
	"github.com/walle692/D0018E/BackEnd/version2/utils"
)

func TestGetUserBasket(t *testing.T) {
	ctx, pool := test_setup.SetUpDB(t)

	sellerID := test_setup.SeedUser(t, ctx, pool, "seller_"+Suffix(), "seller")
	price_1, price_2 := 100.0, 200.0
	p1 := test_setup.SeedProduct(t, ctx, pool, sellerID, price_1, 100)
	p2 := test_setup.SeedProduct(t, ctx, pool, sellerID, price_2, 200)
	t.Run("no items in basket", func(t *testing.T) {
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		_ = test_setup.SeedBasket(t, ctx, pool, buyerID)
		out, err := utils.GetUserBasket(buyerID)
		require.NoError(t, err)
		require.NotEmpty(t, out)
	})

	t.Run("1 item in basket", func(t *testing.T) {
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		basketID := test_setup.SeedBasket(t, ctx, pool, buyerID)
		_ = test_setup.SeedBasketItem(t, ctx, pool, basketID, p1, 1)

		out, err := utils.GetUserBasket(buyerID)
		require.NoError(t, err)
		require.Equal(t, out.Basket_items[0].Product_id, p1)
		require.Equal(t, out.Basket_items[0].Price, price_1)

	})

	t.Run("2 item in basket", func(t *testing.T) {
		buyerID := test_setup.SeedUser(t, ctx, pool, "buyer_"+Suffix(), "buyer")
		basketID := test_setup.SeedBasket(t, ctx, pool, buyerID)
		_ = test_setup.SeedBasketItem(t, ctx, pool, basketID, p1, 1)
		_ = test_setup.SeedBasketItem(t, ctx, pool, basketID, p2, 1)

		out, err := utils.GetUserBasket(buyerID)
		require.NoError(t, err)
		require.Equal(t, out.Basket_items[1].Product_id, p2)
		require.Equal(t, out.Totalprice, price_1+price_2)

	})
}
