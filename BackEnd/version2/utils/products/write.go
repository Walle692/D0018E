package products

import (
	"context"
	"fmt"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func Create(sellerID int, productName string, manufacturer string,
	desc string, pictureURL string, screen_size float64, price float64, stock int) error {

	if screen_size < 0 || price < 0 || stock < 0 {
		return fmt.Errorf("Can't be negative")
	}

	ctx := context.Background()
	pool := global.Get().Pool()

	if _, err := pool.Exec(ctx, `
		INSERT INTO myschema.products(product_name, manufacturer, seller_user_id, 
			description, screen_size, picture_url, price, stock)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`, productName, manufacturer, sellerID, desc, screen_size, pictureURL, price, stock,
	); err != nil {
		return err
	}
	return nil
}

func Delist(productID int) error {
	ctx := context.Background()
	pool := global.Get().Pool()

	if _, err := pool.Exec(ctx, `
	UPDATE myschema.products
	SET active = false
	WHERE product_id = $1
	`, productID); err != nil {
		return err
	}
	return nil
}

func Update(sellerID int, product_id int, productName string, manufacturer string,
	desc string, pictureURL string, screen_size float64, price float64, stock int) error {
	ctx := context.Background()
	pool := global.Get().Pool()

	if screen_size < 0 || price < 0 || stock < 0 {
		return fmt.Errorf("Can't be negative")
	}

	_, err := pool.Exec(ctx, `
		UPDATE myschema.products
		SET product_name = $1, manufacturer = $2, description = $3, 
			screen_size = $4, picture_url = $5, price = $6, stock = $7
			WHERE product_id = $8 AND seller_user_id = $9
		`,
		productName,
		manufacturer,
		desc,
		screen_size,
		pictureURL,
		screen_size,
		stock,
		product_id,
		sellerID,
	)
	if err != nil {
		return err
	}
	return nil
}
