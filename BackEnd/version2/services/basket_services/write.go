package basket_services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/utils/basket"
)

func CheckOut(c *gin.Context) {
	session := sessions.Default(c)
	userIDStr := session.Get(global.UserID)
	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	orderID, err := basket.Checkout(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Checkout success", "order_id": orderID})

}

func AddToBasket(c *gin.Context) {
	// read input from request body
	var req AddToBasketRequest
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
	if err := basket.CreateItem(userID, req.ProductID, req.Quantity); err != nil {
		fmt.Println("DEBUG: CREATE BASKET ITEM ERROR")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not add item to basket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "item added to basket"})

}

func DeleteFromBasket(c *gin.Context) {
	// read input from request body
	var req DeleteBasketRequest
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
	if err := basket.DeleteBasketItemsWithProduct(userID, req.ProductID); err != nil {
		fmt.Println("DEBUG: CREATE BASKET ITEM ERROR")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "item added to basket"})

}
