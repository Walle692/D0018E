package basket_services

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/utils/basket"
)

// GetBasket returns the contents of the user's basket
// gets the basket id, then looks at all basketitems connected to that basket id
// and returns the quantity and price + product name for each item.
func GetBasket(c *gin.Context) {
	session := sessions.Default(c)
	userIDStr := session.Get(global.UserID)
	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	basketItems, err := basket.GetUserBasket(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users basket"})
	}
	// send the json to the client
	c.JSON(http.StatusOK, basketItems)
}
