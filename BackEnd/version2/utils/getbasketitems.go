package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

type BasketReturn struct {
	Product_id int `json:"product_id"`
	Quantity   int `json:"quantity"`
	Price      int `json:"price"`
}

func GetBasket(c *gin.Context, basketid int) []BasketReturn {
	ctx := context.Background()

	// query done this way to easily be able to get more data if wanted
	query := "SELECT (product_id, quantity, price) FROM myschema.basketitem JOIN myschema.products ON myschema.basketitem.product_id=myschema.product.product_id WHERE basket_id=$1"

	// get the db connection from global
	postgres := global.Get()

	// slice to store items
	items := []BasketReturn{}

	// query the database and scan the result into the struct
	rows, err := postgres.Pool().Query(ctx, query, basketid)
	if err != nil {
		fmt.Println("DEBUG: GET BASKET QUERY ERROR")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get basket items"})
		return nil
	}

	for rows.Next() {
		var item BasketReturn
		if err := rows.Scan(&item.Product_id, &item.Quantity, &item.Price); err != nil {
			fmt.Println("DEBUG: GET BASKET SCAN ERROR")
			fmt.Println(err)
			return nil
		}
		items = append(items, item)
	}

	if err != nil {
		fmt.Println("DEBUG: BASKET OTHER ERROR")
		fmt.Println(err)
		// other error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unexpected error"})
		return nil
	}

	// return the items
	return items
}
