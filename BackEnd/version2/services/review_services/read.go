package review_services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/utils/review"
)

func GetProductReviews(c *gin.Context) {
	productIDstring := c.Param("id")

	productID, err := strconv.Atoi(productIDstring)
	if err != nil {
		fmt.Println("DEBUG: PRODUCT ID CONVERSION ERROR")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	reviews, err := review.GetReviewsByProductID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func GetUserReviews(c *gin.Context) {
	session := sessions.Default(c)
	userIDStr := session.Get("user_id")
	userID, err := strconv.Atoi(userIDStr.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	reviews, err := review.GetReviewsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
