package main

import (
	"shopping-cart/database"
	"shopping-cart/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	database.Connect()

	// Create new Gin router
	r := gin.Default()

	// Register all routes (users, items, cart, orders)
	routes.RegisterRoutes(r)

	// Run server on port 8080
	err := r.Run(":8080")
	if err != nil {
		panic("Failed to start server")
	}
}
