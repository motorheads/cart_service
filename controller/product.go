package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motorheads/cart_service/models"
	// "github.com/motorheads/cart_service/storage"
)

func AddProduct(c *gin.Context) {
	var cart models.Cart

	if err := c.ShouldBindJSON(&cart); err != nil {
		fmt.Println("Error while binding JSON cart (add product) ")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func Checkout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
