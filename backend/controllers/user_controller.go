package controllers

import (
    "net/http"
    "shopping-cart/database"
    "shopping-cart/models"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("my_secret_key")

type LoginInput struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Signup(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save to DB
    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
    var input LoginInput
    var user models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username/password"})
        return
    }

    // Generate JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, _ := token.SignedString(jwtKey)

    user.Token = tokenString
    database.DB.Save(&user)

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
