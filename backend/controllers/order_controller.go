package controllers

import (
    "net/http"
    "shopping-cart/database"
    "shopping-cart/models"
    "github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)
    var cart models.Cart
    if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
        return
    }

    order := models.Order{UserID: userID, CartID: cart.ID}
    database.DB.Create(&order)
    c.JSON(http.StatusOK, gin.H{"message": "Order successful"})
}

func ListOrders(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)
    var orders []models.Order
    database.DB.Where("user_id = ?", userID).Find(&orders)
    c.JSON(http.StatusOK, orders)
}
