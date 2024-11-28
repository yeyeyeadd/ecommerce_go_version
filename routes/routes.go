package routes

import (
	"ecommerce-api/controllers"
	"github.com/gin-gonic/gin"
)

// Initialize Routing
func InitRoutes() *gin.Engine {
	router := gin.Default()

	// User routing
	router.POST("/users/register", controllers.Register)
	router.POST("/users/login", controllers.Login)

	// Product routing
	router.GET("/products", controllers.GetProducts)
	router.POST("/products", controllers.CreateProduct)

	// Order routing
	router.POST("/orders", controllers.CreateOrder)

	// Review routing
	router.POST("/reviews", controllers.CreateReview)

	return router
}
