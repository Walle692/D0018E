package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func GetProducts(c *gin.Context) {
	ctx := context.Background()

	// query
	query := "SELECT * FROM myschema.products"

	// get db connection
	postgres := global.Get()

	// make query, all rows are stored in rows
	rows, err := postgres.Pool().Query(ctx, query)
	if err != nil {
		fmt.Println("DEBUG: PRODUCTS QUERY ERROR")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	// slice to store the products
	products := []global.ProductStruct{}

	// iterate trough all rows
	for rows.Next() {
		var p global.ProductStruct
		if err := rows.Scan(&p.Product_id, &p.Product_name, &p.Manufacturer, &p.Seller_user_id, &p.Description, &p.Screen_size, &p.Picture_url, &p.Price, &p.Stock, &p.Active); err != nil {
			fmt.Println("DEBUG: PRODUCTS SCAN ERROR")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("DEBUG: PRODUCTS POST SCAN ERROR")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
	}

	c.JSON(http.StatusOK, products)

}
