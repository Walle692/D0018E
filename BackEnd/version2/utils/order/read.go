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

	queryOrder := head + mid + tail

	queryOrderItem := `
		SELECT oi.quantity, oi.price,
		p.product_id, p.product_name, p.manufacturer, p.picture_url
		FROM myschema.orderitem oi
		JOIN myschema.products p ON oi.product_id = p.product_id
		WHERE oi.order_id = $1
	`
	return orderHelper(queryOrder, []any{searchObject}, queryOrderItem, func(orderID int) []any { return []any{orderID} }, limit, offset)
}

// Build query on on how to fetch orders
// And lastly orderItem query that only returns items that the seller owns
func ListBySellerSearch(sellerID int, searchType string, searchObject any, limit, offset int) ([]Order, error) {
	head := `
		SELECT DISTINCT o.order_id, o.order_user_id, o.orderdate, o.totalprice
		FROM myschema.order o
		JOIN myschema.orderitem oi ON o.order_id = oi.order_id
		JOIN myschema.products p ON oi.product_id = p.product_id
	`
	var mid string
	tail := `
		ORDER BY o.orderdate DESC, o.order_id DESC
		LIMIT $3 OFFSET $4
	`

	switch searchType {
	case "product_id":
		mid = `WHERE oi.product_id = $1 AND p.seller_user_id = $2`
	case "product_name":
		mid = `WHERE p.product_name = $1 AND p.seller_user_id = $2`
	default:
		return nil, fmt.Errorf("Search type not supported")
	}

	queryOrder := head + mid + tail

	queryOrderItem := `
		SELECT oi.quantity, oi.price,
		p.product_id, p.product_name, p.manufacturer, p.picture_url
		FROM myschema.orderitem oi
		JOIN myschema.products p ON oi.product_id = p.product_id
		WHERE oi.order_id = $1 AND p.seller_user_id = $2
	`

	return orderHelper(queryOrder, []any{searchObject, sellerID}, queryOrderItem, func(orderID int) []any { return []any{orderID, sellerID} }, limit, offset)
}

// Builds what is necessary to show a complete order
// 1st pass in query to select which orders to return + array of argument for query
// 2nd pass in query to select what the orders should contain + func(orderID int) []any {return []any{orderID + (query args)}}
// Which really just is a wrapper for []any(query args)
func orderHelper(queryOrder string, orderArgs []any, queryItem string, itemArgsBuilder func(orderID int) []any, limit, offset int) ([]Order, error) {
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
	args := append(append([]any{}, orderArgs...), limit, offset)
	orders, err := queries.ListByQuery(queryOrder, scanOrder, args...)
	if err != nil {
		return nil, err
	}
	for i := range orders {
		itemArgs := itemArgsBuilder(orders[i].Order_id)
		orderItems, err := queries.ListByQuery(queryItem, scanOrderItem, itemArgs...)
		if err != nil {
			return nil, err
		}
		orders[i].Order_items = orderItems

	}
	return orders, nil
}
