package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/motorheads/cart_service/controller"
)

func New() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	api.Use()
	{
		api.POST("/cart", controller.AddProduct)
		api.DELETE("/cart", controller.Checkout)
	}
	return router
}
