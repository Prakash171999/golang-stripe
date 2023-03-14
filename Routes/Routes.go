package Routes

import (
	"proj-mido/stripe-gateway/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	mido := r.Group("/products")
	{
		mido.GET("", Controllers.GetProducts)
		mido.POST("", Controllers.CreateProducts)
		// mido.GET("products/:id", Controllers.GetProductsByID)
		// mido.PUT("products/:id", Controllers.UpdateProducts)
		// mido.DELETE("products/:id", Controllers.DeleteProducts)
	}
	return r
}
