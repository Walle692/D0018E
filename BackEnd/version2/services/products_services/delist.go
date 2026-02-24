package products_services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/utils/products"
)

type DeleteProduct struct {
	Product_id int `json:"product_id"`
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
