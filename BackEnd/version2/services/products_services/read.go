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

func ListAll(c *gin.Context) {
	products, err := products.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetByProductID(c *gin.Context) {
	// string
	productIDsting := c.Param("id")

	// convert string to int
	productID, err := strconv.Atoi(productIDsting)
	if err != nil {
		fmt.Println("DEBUG: PRODUCT ID CONVERSION ERROR")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	p, err := products.GetByProductID(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func GetSellerProduct(c *gin.Context) {
	session := sessions.Default(c)
	userIDStr := session.Get(global.UserID)
	userID, err := strconv.Atoi(userIDStr.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product, err := products.ListBySellerID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
