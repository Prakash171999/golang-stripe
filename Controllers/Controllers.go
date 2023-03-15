package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proj-mido/stripe-gateway/Models"
	"proj-mido/stripe-gateway/Repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func CreateProducts(c *gin.Context) {
	var product Models.Products
	c.BindJSON(&product)
	err := Repository.CreateProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func Config(c *gin.Context) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("FSDFDSDSFDSFDSF", os.Getenv("STRIPE_PUBLISHABLE_KEY"))

	c.JSON(http.StatusOK, gin.H{
		"publishableKey": os.Getenv("STRIPE_PUBLISHABLE_KEY"),
	})
}
