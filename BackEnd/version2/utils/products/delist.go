package products

import (
	"context"

	"github.com/walle692/D0018E/BackEnd/version2/global"
)

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
