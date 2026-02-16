package utils

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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

func GetUserOrders(userID int) ([]OrderResponse, error) {

	ctx := context.Background()
	pool := global.Get().Pool()

	orderRoot, err := fetchOrders(ctx, pool, userID)
	if err != nil {
		return nil, err
	}

	return orderRoot, nil
}

func fetchOrders(ctx context.Context, pool *pgxpool.Pool, userID int) ([]OrderResponse, error) {
	orderIdQuery := `--sql
		SELECT order_id, orderdate, totalprice 
		FROM myschema.order 
		WHERE order_user_id = $1
	`

	orderRoot := make([]OrderResponse, 0)

	rows, err := pool.Query(ctx, orderIdQuery, userID)
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

func fetchOrderItems(ctx context.Context, pool *pgxpool.Pool, orderID int) {
	orderItemQuery := `--sql
			SELECT product_id, quantity, price
			FROM myschema.orderitem
			WHERE order_id = $1
	`
	_ = orderItemQuery
	return
}
