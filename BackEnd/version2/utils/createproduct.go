package utils

import (
	"context"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func CreateProduct(sellerID int, productName string, manufacturer string,
	desc string, pictureURL string, screen_size int, price float64, stock int) error {

	ctx := context.Background()
	pool := global.Get().Pool()

	if _, err := pool.Exec(ctx, `
		INSERT INTO myschema.products(product_name, manufacturer, seller_user_id, 
			description, screen_size, picture_url, price, stock)
		VALUES ($1, $2, $3, $4, $5, $6, $7, 8)
		`, productName, manufacturer, sellerID, desc, screen_size, pictureURL, price, stock,
	); err != nil {
		return err
	}
	return nil
}
