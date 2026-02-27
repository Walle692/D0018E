package products

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/walle692/D0018E/BackEnd/version2/global"
	queries "github.com/walle692/D0018E/BackEnd/version2/utils"
)

func ListBySearch(searchType, searchObject, sortType, sortDirection string, limit, offset int) ([]Product, error) {
	head := `
		SELECT product_id, product_name, manufacturer, description, screen_size, picture_url, price, stock
		FROM myschema.products p
		`
	var search, s2, s3 string
	tail := `--sql
		LIMIT $2 OFFSET $3
	`

	switch searchType {
	case "product_name":
		search = `--sql
			WHERE p.product_name = $1 AND p.active = true	
		`

	case "manufacturer":
		search = `--sql
			WHERE p.manufacturer = $1 AND p.active = true	
		`
	default:
		search = `--sql
			WHERE p.active = true AND $1::text IS NOT NULL
		`
	}

	switch sortType {
	case "price":
		s2 = `--sql
			ORDER BY price
		`
	case "screen_size":
		s2 = `--sql
			ORDER BY screen_size
		`
	default:
		return nil, fmt.Errorf("Invalid sort type")
	}

	switch sortDirection {
	case "descending":
		s3 = `DESC`
	default:
		s3 = ``
	}

	query := head + search + s2 + s3 + tail

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

	return queries.ListByQuery(query, scanProduct, searchObject, limit, offset)
}

func ListAll() ([]Product, error) {

	query := `
		SELECT product_id, product_name, manufacturer, description, screen_size, picture_url, price, stock, active 
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
			&p.Active,
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
		SELECT product_id, product_name, manufacturer, description, screen_size, picture_url, price, stock, active  
		FROM myschema.products
		WHERE seller_user_id = $1 AND active = true
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
			&p.Active,
		)
		return p, err
	}

	p, err := queries.ListByQuery(query, scanProduct, sellerID)
	if err != nil {
		return nil, err
	}
	return p, nil
}
