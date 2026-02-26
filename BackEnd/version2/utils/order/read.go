package order

import (
	"github.com/jackc/pgx/v5"
	queries "github.com/walle692/D0018E/BackEnd/version2/utils"
)

func ListByUserID(userID, limit, offset int) ([]Order, error) {

	if limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	queryOrder := `
		SELECT order_id, order_user_id, orderdate, totalprice
		FROM myschema.order
		WHERE order_user_id = $1
		ORDER BY orderdate DESC, order_id DESC
		LIMIT $2 OFFSET $3
	`

	scanOrder := func(rows pgx.Rows) (Order, error) {
		var o Order
		err := rows.Scan(
			&o.Order_id,
			&o.Order_user_id,
			&o.Order_date,
			&o.Total_price,
		)
		return o, err
	}

	queryOrderItem := `
		SELECT oi.quantity, oi.price,
		p.product_id, p.product_name, p.manufacturer, p.picture_url
		FROM myschema.orderitem oi
		JOIN myschema.products p ON oi.product_id = p.product_id
		WHERE oi.order_id = $1
	`
	scanOrderItem := func(rows pgx.Rows) (OrderItem, error) {
		var o OrderItem
		err := rows.Scan(
			&o.Quantity,
			&o.Price,
			&o.Product_id,
			&o.Product_name,
			&o.Manufacturer,
			&o.Picture_url,
		)
		return o, err
	}
	orders, err := queries.ListByQuery(queryOrder, scanOrder, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	for i := range orders {

		orderItems, err := queries.ListByQuery(queryOrderItem, scanOrderItem, orders[i].Order_id)
		if err != nil {
			return nil, err
		}
		orders[i].Order_items = orderItems

	}
	return orders, nil

}
