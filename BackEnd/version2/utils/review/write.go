package review

import (
	"context"
	"errors"
	"fmt"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func WriteReview(customer_user_id, product_id int, comment string, rating int) error {
	ctx := context.Background()
	pool := global.Get().Pool()

	// check rating range
	if rating < 1 || rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	// check if product exists
	var exists bool
	err := pool.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM myschema.products WHERE product_id=$1)`, product_id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("check product exists: %w", err)
	}

	if !exists {
		return errors.New("product not found")

	}

	// check if the user has already reviewed the product
	err = pool.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM myschema.rating WHERE customer_user_id=$1 AND product_id=$2)`, customer_user_id, product_id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("check existing review: %w", err)
	}
	if exists {
		return errors.New("user has already reviewed this product")
	}

	// insert review
	_, err = pool.Exec(ctx, `
		INSERT INTO myschema.rating (customer_user_id, product_id, comment, rating) 
		VALUES ($1, $2, $3, $4)
	`, customer_user_id, product_id, comment, rating)
	if err != nil {
		return fmt.Errorf("insert review: %w", err)
	}

	return nil
}
