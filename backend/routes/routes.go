package routes

import (
    "github.com/gin-gonic/gin"
    "shopping-cart/controllers"
    "shopping-cart/middleware"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/users", controllers.Signup)
    r.POST("/users/login", controllers.Login)

    r.POST("/items", controllers.CreateItem)
    r.GET("/items", controllers.ListItems)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.POST("/carts", controllers.AddToCart)
        auth.GET("/carts", controllers.ListCart)
        auth.POST("/orders", controllers.PlaceOrder)
        auth.GET("/orders", controllers.ListOrders)
    }
}
