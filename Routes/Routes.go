package Routes

import (
	"proj-mido/stripe-gateway/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	mido := r.Group("/mido")
	{
		mido.GET("products", Controllers.GetProducts)
		mido.POST("products", Controllers.CreateProducts)
		mido.GET("config", Controllers.Config)
		mido.POST("create-payment-intent", Controllers.HandleCreatePaymentIntent)
		// mido.GET("products/:id", Controllers.GetProductsByID)
		// mido.PUT("products/:id", Controllers.UpdateProducts)
		// mido.DELETE("products/:id", Controllers.DeleteProducts)

	}
	return r
}
