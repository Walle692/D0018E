package queries

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

/*
Helper function to get multiple rows from any query
Input query string to be executed, and function name of the scaning method example:

	func (){
		query := `
			SELECT product_name
			FROM myschema.products
			WHERE seller_user_id = $1 AND active = true
		`

		scanProduct := func(rows pgx.Rows) (Product, error) {
			var p Product
			err := rows.Scan(
				&p.Product_name,
			)
			return p, err
		}
	}
	p, err := queries.ListByQuery(query, scanProduct, args)
*/
func ListByQuery[T any](query string, scanRow func(rows pgx.Rows) (T, error), args ...any) ([]T, error) {
	ctx := context.Background()
	pool := global.Get().Pool()

	items := make([]T, 0)

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item, err := scanRow(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
