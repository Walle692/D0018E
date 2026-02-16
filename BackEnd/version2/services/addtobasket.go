package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/utils"
)

// takes the product id and quantity from the request and adds the product to the user's basket
type AddToBasketRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,min=1"`
}

func AddToBasket(c *gin.Context) {
	// read input from request body
	var req AddToBasketRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("DEBUG: AddToBasketRequest: %+v\n", req)

	// get the users basket id
	basketID, err := utils.GetBasketID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get basket"})
		return
	}

	// create the basketItem in the database
	if err := utils.CreateBasketItem(c, basketID, req.ProductID, req.Quantity); err != nil {
		fmt.Println("DEBUG: CREATE BASKET ITEM ERROR")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not add item to basket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "item added to basket"})

}
