package controllers

import (
	"ecommerce-api/models"
	"ecommerce-api/utils"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var input struct {
		Items []struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		} `json:"items"`
	}
	// Bind request parameters.
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Extract the token from the Authorization header of the request.
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	// Extract the user ID from the token.
	token = strings.TrimPrefix(token, "Bearer ")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Assume the token contains a "user_id" field that stores the user's ID.
	buyerID := claims["user_id"].(float64)

	// Create Order
	order := models.Order{
		BuyerID: uint(buyerID),
	}

	if err := models.DB.Create(&order).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		log.Println("Error creating order:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Add product for order
	for _, item := range input.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		if err := models.DB.Create(&orderItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add order items"})
			return
		}
	}

	// Return success
	c.JSON(http.StatusCreated, gin.H{
		"id":       order.ID,
		"buyer_id": order.BuyerID,
		"items":    input.Items,
	})
}
