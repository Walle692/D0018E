package products

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/walle692/D0018E/BackEnd/version2/global"
	queries "github.com/walle692/D0018E/BackEnd/version2/utils"
)

func ListAll() ([]Product, error) {

	query := `
		SELECT product_id, product_name, manufacturer, description, screen_size, picture_url, price, stock  
		FROM myschema.products
		WHERE active = true
	`
	scanProduct := func(rows pgx.Rows) (Product, error) {
		var p Product
		err := rows.Scan(
			&p.Product_id,
			&p.Product_name,
			&p.Manufacturer,
			&p.Description,
			&p.Screen_size,
			&p.Picture_url,
			&p.Price,
			&p.Stock,
		)
		return p, err
	}

	p, err := queries.ListByQuery(query, scanProduct)
	if err != nil {
		return nil, err
	}
	return p, nil
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
	query := `
		SELECT product_name, manufacturer, description, screen_size, picture_url, price, stock  
		FROM myschema.products
		WHERE seller_user_id = $1 AND active = true
	`

	scanProduct := func(rows pgx.Rows) (Product, error) {
		var p Product
		err := rows.Scan(
			&p.Product_name,
			&p.Manufacturer,
			&p.Description,
			&p.Screen_size,
			&p.Picture_url,
			&p.Price,
			&p.Stock,
		)
		return p, err
	}

	p, err := queries.ListByQuery(query, scanProduct, sellerID)
	if err != nil {
		return nil, err
	}
	return p, nil
}
