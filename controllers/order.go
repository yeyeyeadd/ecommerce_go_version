package controllers

import (
	"ecommerce-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create order
func CreateOrder(c *gin.Context) {
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
}
