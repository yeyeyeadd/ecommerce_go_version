package controllers

import (
	"ecommerce-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Add Comment
func CreateReview(c *gin.Context) {
	var input struct {
		UserID    uint   `json:"user_id" binding:"required"`
		ProductID uint   `json:"product_id" binding:"required"`
		OrderID   uint   `json:"order_id" binding:"required"`
		Rating    int    `json:"rating" binding:"required,min=1,max=5"`
		Comment   string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := models.Review{
		UserID:    input.UserID,
		ProductID: input.ProductID,
		OrderID:   input.OrderID,
		Rating:    input.Rating,
		Comment:   input.Comment,
	}
	if err := models.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}
