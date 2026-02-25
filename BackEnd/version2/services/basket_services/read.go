package basket_services

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/utils/basket"
)

func GetBasket(c *gin.Context) {
	session := sessions.Default(c)
	userIDStr := session.Get(global.UserID)
	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	basketItems, err := basket.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users basket"})
	}
	// send the json to the client
	c.JSON(http.StatusOK, basketItems)
}
