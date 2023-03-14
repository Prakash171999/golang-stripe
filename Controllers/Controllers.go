package Controllers

import (
	"net/http"
	"proj-mido/stripe-gateway/Models"
	"proj-mido/stripe-gateway/Repository"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []Models.Products
	err := Repository.GetAllProducts(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}
