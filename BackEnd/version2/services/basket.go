package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/utils"
)

// GetBasket returns the contents of the user's basket
// gets the basket id, then looks at all basketitems connected to that basket id
// and returns the quantity and price + product name for each item.
func GetBasket(c *gin.Context) {

	// get the users basket id
	basketID, err := utils.GetBasketID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get basket"})
		return
	}

	// get the basket items for the basket id
	items := utils.GetBasket(c, basketID)
	if items == nil {
		// error already handled in utils.GetBasket
		return
	}

	// send the json to the client
	c.JSON(http.StatusOK, items)
}
