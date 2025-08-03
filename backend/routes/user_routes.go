package routes

import (
    "github.com/gin-gonic/gin"
    "shopping-cart/controllers"
)

func UserRoutes(r *gin.Engine) {
    r.POST("/users", controllers.Signup)
    r.POST("/users/login", controllers.Login)
}
