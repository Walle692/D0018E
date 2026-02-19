package utils

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func CreateBasketItem(c *gin.Context, basketID int, productID int, quantity int) error {
	ctx := context.Background()

	postgres := global.Get()

	/*
		// start by getting the price of the product from the database
		var price int
		query := "SELECT price FROM myschema.products WHERE product_id=$1"
		err := postgres.Pool().QueryRow(c, query, productID).Scan(&price)
		if err != nil {
			fmt.Println("DEBUG: CREATE BASKET ITEM GET PRICE ERROR")
			return err
		}
	*/

	// insert the basketitem into the db
	query := "INSERT INTO myschema.basketitem (basket_id, product_id, quantity) VALUES ($1, $2, $3, $4)"
	_, err := postgres.Pool().Exec(ctx, query, basketID, productID, quantity)
	if err != nil {
		fmt.Println("DEBUG: CREATE BASKET ITEM INSERT ERROR")
		return err
	}

	return nil
}
