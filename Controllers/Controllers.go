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
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
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

func HandleCreatePaymentIntent(c *gin.Context) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	type item struct {
		id string
	}

	var req struct {
		Items []item `json:"items"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("ShouldBindJSON: %v", err)
		return
	}

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		// Amount:   stripe.Int64(calculateOrderAmount(req.Items)),
		Amount:   stripe.Int64(1400),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("pi.New: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"clientSecret": pi.ClientSecret,
	})
}
