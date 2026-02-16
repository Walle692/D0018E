package utils

import (
	"context"
	"time"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

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

type OrderResponse struct {
	Order_id    int                    `json:"order_id"`
	Order_items []orderProductResponse `json:"order_items"`
	Order_date  time.Time              `json:"orderdate"`
	Total_price float64                `json:"totalprice"`
}

func GetUserOrders(userId int) ([]OrderResponse, error) {

	ctx := context.Background()
	postgres := global.Get()

	orderIdQuery := `--sql
		SELECT order_id, orderdate, totalprice 
		FROM myschema.order 
		WHERE order_user_id = $1
	`

	orderRoot := make([]OrderResponse, 0)

	rows, err := postgres.Pool().Query(ctx, orderIdQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o OrderResponse
		if err := rows.Scan(&o.Order_id, &o.Order_date, &o.Total_price); err != nil {
			return nil, err
		}
		orderRoot = append(orderRoot, o)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orderRoot, nil
}
