package utils

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func GetProduct(c *gin.Context) {
	ctx := context.Background()

	// string
	productIDsting := c.Param("id")

	// convert string to int
	productID, err := strconv.Atoi(productIDsting)
	if err != nil {
		fmt.Println("DEBUG: PRODUCT ID CONVERSION ERROR")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// create struct for the product
	var p global.ProductStruct

	// query
	query := "SELECT * FROM myschema.products WHERE product_id=$1"

	// get the db connection from global
	postgres := global.Get()

	// query the database and scan the result into the struct
	err = postgres.Pool().QueryRow(ctx, query, productID).Scan(&p.Product_id, &p.Product_name, &p.Manufacturer, &p.Seller_user_id, &p.Description, &p.Screen_size, &p.Picture_url, &p.Price, &p.Stock, &p.Active)

	if err == pgx.ErrNoRows {
		fmt.Println("DEBUG: PRODUCT ERR NO ROWS")
		// no product found
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	} else if err != nil {
		fmt.Println("DEBUG: PRODUCT OTHER GET ERROR")
		fmt.Println(err)
		// other error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unexpected error"})
		return
	}

	// send the json to the client
	c.JSON(http.StatusOK, p)
}
