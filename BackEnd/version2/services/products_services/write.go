package products_services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/utils/products"
)

func CreateProduct(c *gin.Context) {
	// read input from request body
	var req CreateProductRequest
	session := sessions.Default(c)
	userIDStr := session.Get(global.UserID)
	userID, err := strconv.Atoi(userIDStr.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create the basketItem in the database
	if err := products.CreateProduct(userID, req.ProductName, req.Manufacturer, req.Description, req.PictureURL, req.ScreenSize, req.Price, req.Stock); err != nil {
		fmt.Println("DEBUG: CREATE BASKET ITEM ERROR")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "item added to basket"})

}

func Delist(c *gin.Context) {
	var req DeleteProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("here")
		return
	}

	// create the basketItem in the database
	if err := products.Delist(req.Product_id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete product"})
		log.Printf("New er")
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "product delisted"})

}
