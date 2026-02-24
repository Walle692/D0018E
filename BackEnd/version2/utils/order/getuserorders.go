package order

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

type OrderItemResponse struct {
	Product_id   int    `json:"product_id"`
	Product_name string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
	Picture_url  string `json:"picture_url"`
}

type OrderProductResponse struct {
	Product  OrderItemResponse `json:"product"`
	Quantity int               `json:"quantity"`
	Price    float32           `json:"price"`
}

type OrderResponse struct {
	Order_id    int                    `json:"order_id"`
	Order_items []OrderProductResponse `json:"order_items"`
	Order_date  time.Time              `json:"orderdate"`
	Total_price float64                `json:"totalprice"`
}

func GetUserOrders(userID int) ([]OrderResponse, error) {

	ctx := context.Background()
	pool := global.Get().Pool()

	orderRoot, orderIDs, err := fetchOrders(ctx, pool, userID)
	if err != nil {
		return nil, err
	}
	//Map how add into order_items
	idx := map[int]int{}
	for i := range orderRoot {
		idx[orderRoot[i].Order_id] = i
	}

	itemsByOrder, err := fetchOrderItems(ctx, pool, orderIDs)
	if err != nil {
		return nil, err
	}

	// attach
	for orderID, items := range itemsByOrder {
		i := idx[orderID]
		orderRoot[i].Order_items = append(orderRoot[i].Order_items, items...)
	}

	return orderRoot, nil
}

func fetchOrders(ctx context.Context, pool *pgxpool.Pool, userID int) ([]OrderResponse, []int, error) {
	orderIdQuery := `--sql
		SELECT order_id, orderdate, totalprice 
		FROM myschema.order 
		WHERE order_user_id = $1
	`

	orderRoot, orderID := make([]OrderResponse, 0), make([]int, 0)

	rows, err := pool.Query(ctx, orderIdQuery, userID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var o OrderResponse
		if err := rows.Scan(&o.Order_id, &o.Order_date, &o.Total_price); err != nil {
			return nil, nil, err
		}
		orderRoot = append(orderRoot, o)
		orderID = append(orderID, o.Order_id)
	}
	if err = rows.Err(); err != nil {
		return nil, nil, err
	}
	return orderRoot, orderID, nil
}

func fetchOrderItems(ctx context.Context, pool *pgxpool.Pool, orderIDs []int) (map[int][]OrderProductResponse, error) {
	orderItemQuery := `--sql
    SELECT
    oi.order_id, oi.quantity, oi.price,
    p.product_id, p.product_name, p.manufacturer, p.picture_url
    FROM myschema.orderitem oi
    JOIN myschema.products p ON p.product_id = oi.product_id
    WHERE oi.order_id = ANY($1)
    ORDER BY oi.order_id, oi.order_item_id
	`

	rows, err := pool.Query(ctx, orderItemQuery, orderIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make(map[int][]OrderProductResponse)

	for rows.Next() {
		var orderID int
		var qty int
		var itemPrice float32
		var p OrderItemResponse
		if err := rows.Scan(&orderID, &qty, &itemPrice, &p.Product_id, &p.Product_name, &p.Manufacturer, &p.Picture_url); err != nil {
			return nil, err
		}
		item := OrderProductResponse{
			Product:  p,
			Quantity: qty,
			Price:    itemPrice,
		}
		out[orderID] = append(out[orderID], item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return out, nil
}
