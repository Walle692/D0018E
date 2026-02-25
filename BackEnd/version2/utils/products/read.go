package products

import (
	"context"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func ListAll() ([]Product, error) {
	ctx := context.Background()
	pool := global.Get().Pool()

	query := `SELECT * FROM myschema.products WHERE active=true`

	products := make([]Product, 0)

	rows, err := pool.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Product_id, &p.Product_name, &p.Manufacturer, &p.Seller_user_id, &p.Description,
			&p.Screen_size, &p.Picture_url, &p.Price, &p.Stock, &p.Active); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetByProductID(productID int) (Product, error) {
	ctx := context.Background()
	pool := global.Get().Pool()

	query := `SELECT * FROM myschema.products WHERE product_id=$1`

	var p Product

	err := pool.QueryRow(ctx, query, productID).Scan(&p.Product_id, &p.Product_name, &p.Manufacturer, &p.Seller_user_id, &p.Description,
		&p.Screen_size, &p.Picture_url, &p.Price, &p.Stock, &p.Active)
	if err != nil {
		return p, err
	}
	return p, nil
}

func ListBySellerID(sellerID int) ([]Product, error) {
	ctx := context.Background()
	pool := global.Get().Pool()

	products := make([]Product, 0)

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
		var s Product
		if err := rows.Scan(&s.Product_id, &s.Product_name, &s.Picture_url, &s.Price, &s.Stock,
			&s.Active, &s.Screen_size, &s.Description, &s.Manufacturer); err != nil {
			return nil, err
		}
		products = append(products, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
