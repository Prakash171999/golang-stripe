package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proj-mido/stripe-gateway/Models"
	"proj-mido/stripe-gateway/Repository"
	"strconv"

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
		c.JSON(http.StatusOK, gin.H{"products": products})
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

	c.JSON(http.StatusOK, gin.H{
		"publishableKey": os.Getenv("STRIPE_PUBLISHABLE_KEY"),
	})
}

func HandleCreatePaymentIntent(c *gin.Context) {

	var product Models.Products

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("ShouldBindJSON: %v", err)
		return
	}

	product_id := strconv.FormatUint(uint64(product.Id), 10)

	data, err := Repository.GetAProduct(&product, product_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	fmt.Println("daaaaata==>", data)

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(data.Price)),
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
