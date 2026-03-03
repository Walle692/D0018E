package review

import (
	"github.com/jackc/pgx/v5"
	queries "github.com/walle692/D0018E/BackEnd/version2/utils"
)

func GetReviewsByProductID(productID int) ([]Review, error) {
	query := `
		SELECT comment_id, customer_user_id, product_id, comment, rating 
		FROM myschema.rating
		WHERE product_id = $1
	`

	scanRewiew := func(rows pgx.Rows) (Review, error) {
		var r Review
		err := rows.Scan(&r.Comment_id, &r.Customer_user_id, &r.Product_id, &r.Comment, &r.Rating)
		return r, err
	}

	r, err := queries.ListByQuery(query, scanRewiew, productID)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetReviewsByUserID(userID int) ([]Review, error) {
	query := `
		SELECT comment_id, product_id, comment, rating 
		FROM myschema.rating
		WHERE customer_user_id = $1
	`

	scanRewiew := func(rows pgx.Rows) (Review, error) {
		var r Review
		err := rows.Scan(&r.Comment_id, &r.Product_id, &r.Comment, &r.Rating)
		return r, err
	}

	r, err := queries.ListByQuery(query, scanRewiew, userID)
	if err != nil {
		return nil, err
	}

	return r, nil
}
