package order

import "time"

type Order struct {
	Order_id      int         `json:"order_id"`
	Order_user_id int         `json:"order_user_id"`
	Order_date    time.Time   `json:"orderdate"`
	Total_price   float64     `json:"totalprice"`
	Order_items   []OrderItem `json:"order_items"`
}

type OrderItem struct {
	// From order item
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	// From product
	Product_id   int    `json:"product_id"`
	Product_name string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
	Picture_url  string `json:"picture_url"`
}
