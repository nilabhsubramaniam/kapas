package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCart returns user's cart
func GetCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get cart endpoint - To be implemented",
	})
}

// AddToCart adds an item to cart
func AddToCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Add to cart endpoint - To be implemented",
	})
}

// UpdateCartItem updates cart item quantity
func UpdateCartItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update cart item endpoint - To be implemented",
	})
}

// RemoveFromCart removes an item from cart
func RemoveFromCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Remove from cart endpoint - To be implemented",
	})
}

// ClearCart clears all cart items
func ClearCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Clear cart endpoint - To be implemented",
	})
}
