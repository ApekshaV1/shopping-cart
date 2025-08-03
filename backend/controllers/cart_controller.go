package controllers

import (
    "net/http"
    "shopping-cart/database"
    "shopping-cart/models"
    "github.com/gin-gonic/gin"
)

type AddToCartInput struct {
    ItemID uint `json:"item_id"`
}

func AddToCart(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)
    var cart models.Cart

    database.DB.Where("user_id = ?", userID).FirstOrCreate(&cart, models.Cart{UserID: userID})

    var input AddToCartInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    cartItem := models.CartItem{CartID: cart.ID, ItemID: input.ItemID}
    database.DB.Create(&cartItem)

    c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func ListCart(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)
    var cart models.Cart
    if err := database.DB.Preload("Items").Where("user_id = ?", userID).First(&cart).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
        return
    }

    c.JSON(http.StatusOK, cart.Items)
}
