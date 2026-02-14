package utils_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/walle692/D0018E/BackEnd/version2/test_setup"
	"github.com/walle692/D0018E/BackEnd/version2/utils"
)

//Can return multiple orders

type orderItemResponse struct {
	Product_id     int     `json:"product_id"`
	Product_name   string  `json:"product_name"`
	Manufacturer   string  `json:"manufacturer"`
	Seller_user_id int     `json:"seller_user_id"`
	Description    string  `json:"description"`
	Screen_size    float32 `json:"screen_size"`
	Picture_url    string  `json:"picture_url"`
	Sku            string  `json:"sku"`
	Price          float32 `json:"price"`
	Stock          int     `json:"stock"`
}

type orderProductResponse struct {
	Product  []orderItemResponse `json:"product"`
	Quantity int                 `json:"quantity"`
	Price    float32             `json:"price"`
}

type orderResponse struct {
	Order_items []orderProductResponse `json:"order_items"`
	Order_date  string                 `json:"orderdate"`
	Total_price float64                `json:"totalprice"`
}

func TestGetUserBasket(t *testing.T) {

	ctx, pool := test_setup.SetUpDB(t)
	// Seed buyer users
	bUserTest := "buyerTest_" + fmt.Sprint(time.Now().UnixNano())
	sUserTest := "sellerTest_" + fmt.Sprint(time.Now().UnixNano())

	buyerID := test_setup.SeedUser(t, ctx, pool, bUserTest, "buyer")
	sellerID := test_setup.SeedUser(t, ctx, pool, sUserTest, "seller")

	// Seed product
	productID := test_setup.SeedProduct(t, ctx, pool, sellerID, 100, 10)

	// Seed Order
	orderID := test_setup.SeedOrder(t, ctx, pool, buyerID, "now", 300)
	_ = test_setup.SeedOrderItem(t, ctx, pool, orderID, productID, 3, 900)
	// Seed OrderItem
	// Seeding done now testing function

	out, err := utils.GetUserOrders(nil, buyerID)
	require.NoError(t, err)
	//Check json unpacking
	var orders []orderResponse
	err = json.Unmarshal(out, &orders)
	require.NoError(t, err)

}
