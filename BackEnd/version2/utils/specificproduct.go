package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func GetSpecificProduct(c *gin.Context, productID int) ([]byte, error) {
	ctx := context.Background()

	var p global.ProductStruct

	// query
	query := "SELECT * FROM myschema.products WHERE product_id=$1"

	// get the db connection from global
	postgres := global.Get()

	// query the database and select the password from the username and store in password
	err := postgres.Pool().QueryRow(ctx, query, productID).Scan(&p.Product_id, &p.Product_name, &p.Manufacturer, &p.Seller_user_id, &p.Description, &p.Screen_size, &p.Picture_url, &p.Sku, &p.Price, &p.Stock)

	if err == pgx.ErrNoRows {
		fmt.Println("DEBUG: PRODUCT ERR NO ROWS")
		// no user found
		return make([]byte, 0), errors.New("No user found")
	} else if err != nil {
		fmt.Println("DEBUG: PRODUCT OTHER GET ERROR")
		fmt.Println(err)
		// other error
		return make([]byte, 0), errors.New("Unhandled product error")
	}

	// make the struct into acutal json
	productJSON, err := json.Marshal(p)
	if err != nil {
		fmt.Println("DEBUG: PRODUCT JSON MARSHAL ERROR")
		fmt.Println(err)
		return make([]byte, 0), errors.New("couldn't marshal json")
	}

	return productJSON, nil
}
