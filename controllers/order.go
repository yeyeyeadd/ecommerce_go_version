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

// Create order
/*func CreateOrder(c *gin.Context) {
	var input struct {
		BuyerID uint `json:"buyer_id" binding:"required"`
		Items   []struct {
			ProductID uint `json:"product_id" binding:"required"`
			Quantity  int  `json:"quantity" binding:"required"`
		} `json:"items" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var total float64
	var orderItems []models.OrderItem
	for _, item := range input.Items {
		var product models.Product
		if err := models.DB.First(&product, item.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}
		if product.Stock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
			return
		}

		// Reduce inventory
		product.Stock -= item.Quantity
		models.DB.Save(&product)

		orderItem := models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price * float64(item.Quantity),
		}
		total += orderItem.Price
		orderItems = append(orderItems, orderItem)
	}

	order := models.Order{
		BuyerID:    input.BuyerID,
		TotalPrice: total,
		OrderItems: orderItems,
	}
	if err := models.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}*/

func CreateOrder(c *gin.Context) {
	var input struct {
		Items []struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		} `json:"items"`
	}
	// Binding request parameters
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var total float64
	var orderItems []models.OrderItem
	for _, item := range input.Items {
		var product models.Product
		if err := models.DB.First(&product, item.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}
		if product.Stock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
			return
		}

		// Reduce inventory
		product.Stock -= item.Quantity
		models.DB.Save(&product)

		orderItem := models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price * float64(item.Quantity),
		}
		total += orderItem.Price
		orderItems = append(orderItems, orderItem)
	}

	// Extract the token from the Authorization header of the request
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	// Get the userId from the token
	token = strings.TrimPrefix(token, "Bearer ")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	buyerID := claims["user_id"].(float64)

	order := models.Order{
		BuyerID:    uint(buyerID),
		TotalPrice: total,
		OrderItems: orderItems,
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

	c.JSON(http.StatusCreated, gin.H{
		"id":       order.ID,
		"buyer_id": order.BuyerID,
		"items":    input.Items,
	})
}
