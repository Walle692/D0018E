package basket

import (
	"github.com/jackc/pgx/v5"
	queries "github.com/walle692/D0018E/BackEnd/version2/utils"
)

func GetByUserID(userID int) (BasketResponse, error) {
	var basket BasketResponse

	//Get all basketitems
	query := `--sql
		SELECT bi.quantity, bi.product_id, p.product_name, p.manufacturer, p.picture_url, p.price, p.active
		FROM myschema.basket b
		INNER JOIN myschema.basketitem bi ON b.basket_id = bi.basket_id
		INNER JOIN myschema.products p ON p.product_id = bi.product_id
		WHERE b.basket_user_id = $1
	`

	scan := func(rows pgx.Rows) (BasketItemResponse, error) {
		var bi BasketItemResponse
		err := rows.Scan(
			&bi.Quantity,
			&bi.Product_id,
			&bi.Product_name,
			&bi.Manufacturer,
			&bi.Picture_url,
			&bi.Price,
			&bi.Available,
		)
		return bi, err
	}

	// Calculate total price of items
	items, err := queries.ListByQuery(query, scan, userID)
	if err != nil {
		return basket, err
	}
	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	basket.Basket_items = items
	basket.Totalprice = total

	return basket, nil
}
