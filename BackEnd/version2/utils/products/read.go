package products

import (
	"context"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func GetSellerProduct(sellerID int) ([]SellerProduct, error) {
	ctx := context.Background()
	pool := global.Get().Pool()

	products := make([]SellerProduct, 0)

	rows, err := pool.Query(ctx, `
		SELECT product_id, product_name, picture_url, price, stock, active, screen_size, description, manufacturer
		FROM myschema.products
		WHERE seller_user_id = $1 AND active = true
	`, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var s SellerProduct
		if err := rows.Scan(&s.Product_id, &s.Product_name, &s.Picture_url, &s.Price, &s.Stock, &s.Active, &s.Screen_size, &s.Description, &s.Manufacturer); err != nil {
			return nil, err
		}
		products = append(products, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
