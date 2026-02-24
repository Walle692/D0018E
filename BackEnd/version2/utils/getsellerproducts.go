package utils

import (
	"context"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

type SellerProduct struct {
	Product_id   int     `json:"product_id"`
	Product_name string  `json:"product_name"`
	Manufacturer string  `json:"manufacturer"`
	Picture_url  string  `json:"picture_url"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Active       bool    `json:"active"`
	Description  string  `json:"description"`
	Screen_size  float32 `json:"screen_size"`
}

func GetSellerProduct(sellerID int) ([]SellerProduct, error) {
	ctx := context.Background()
	pool := global.Get().Pool()

	products := make([]SellerProduct, 0)

	rows, err := pool.Query(ctx, `
		SELECT product_id, product_name, picture_url, price, stock, active, screen_size, description, manufacturer
		FROM myschema.products
		WHERE seller_user_id = $1
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
