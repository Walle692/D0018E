package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/utils"
)

func CheckOut(c *gin.Context) {

	basketID, err := utils.GetBasketID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get basket"})
		return
	}

	orderID, err := utils.ConvertBasketToOrder(basketID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not convert basket"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "Checkout success", "order_id": orderID})

}
