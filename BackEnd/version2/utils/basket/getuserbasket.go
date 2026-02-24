package basket

import (
	"context"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

type BasketResponse struct {
	Totalprice   float64              `json:"totalprice"`
	Basket_items []BasketItemResponse `json:"basket_items"`
}

type BasketItemResponse struct {
	Product_id   int     `json:"product_id"`
	Product_name string  `json:"product_name"`
	Manufacturer string  `json:"manufacturer"`
	Picture_url  string  `json:"picture_url"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	Available    bool    `json:"available"`
}

func GetUserBasket(userID int) (BasketResponse, error) {

	ctx := context.Background()
	pool := global.Get().Pool()

	var basket BasketResponse

	//Get all basketitems
	query := `--sql
		SELECT bi.quantity, bi.product_id, p.product_name, p.manufacturer, p.picture_url, p.price, p.active
		FROM myschema.basket b
		INNER JOIN myschema.basketitem bi ON b.basket_id = bi.basket_id
		INNER JOIN myschema.products p ON p.product_id = bi.product_id
		WHERE b.basket_user_id = $1
	`

	rows, err := pool.Query(ctx, query, userID)
	if err != nil {
		return basket, nil
	}
	defer rows.Close()

	basketItem, total := make([]BasketItemResponse, 0), 0.0
	for rows.Next() {
		var bi BasketItemResponse
		if err := rows.Scan(&bi.Quantity, &bi.Product_id, &bi.Product_name, &bi.Manufacturer, &bi.Picture_url, &bi.Price, &bi.Available); err != nil {
			return basket, err
		}
		total = total + bi.Price*float64(bi.Quantity)
		basketItem = append(basketItem, bi)
	}
	if err = rows.Err(); err != nil {
		return basket, err
	}
	basket.Basket_items = basketItem
	basket.Totalprice = total

	return basket, nil
}
