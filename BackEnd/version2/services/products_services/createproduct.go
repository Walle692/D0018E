package products_services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/utils/products"
)

// takes the product id and quantity from the request and adds the product to the user's basket
type CreateProductRequest struct {
	ProductName  string  `json:"product_name" binding:"required"`
	Manufacturer string  `json:"manufacturer" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	ScreenSize   int     `json:"screen_size" binding:"required"`
	PictureURL   string  `json:"picture_url" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Stock        int     `json:"stock" binding:"required"`
}

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
