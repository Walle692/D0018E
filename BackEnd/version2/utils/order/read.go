package order

import (
	"fmt"

	"github.com/jackc/pgx/v5"
	queries "github.com/walle692/D0018E/BackEnd/version2/utils"
)

func ListByUserID(userID, limit, offset int) ([]Order, error) {
	return ListBySearch("userID", userID, limit, offset)
}

// Builds SQL query to fetch rows from order table
// Inputs search term username/userID/orderID
// Returns all orders which matches that search
func ListBySearch(searchType string, searchObject any, limit, offset int) ([]Order, error) {
	head := `
		SELECT o.order_id, o.order_user_id, o.orderdate, o.totalprice
		FROM myschema.order o
	`
	var mid string
	tail := `
		ORDER BY o.orderdate DESC, o.order_id DESC
		LIMIT $2 OFFSET $3
	`
	switch searchType {
	case "username":
		mid = `
			JOIN myschema.users u ON o.order_user_id = u.user_id
			WHERE u.username = $1
		`
	case "orderID":
		mid = `WHERE o.order_id = $1`
	case "userID":
		mid = `
			WHERE o.order_user_id = $1
		`
	default:
		return nil, fmt.Errorf("Invalid search type")
	}

	query := head + mid + tail

	return orderHelper(query, searchObject, limit, offset)
}

// First gets all orders that the query finds
// Then for each returned order adds all items that the each order points to
func orderHelper(query string, input any, limit, offset int) ([]Order, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

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
	orders, err := queries.ListByQuery(query, scanOrder, input, limit, offset)
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
