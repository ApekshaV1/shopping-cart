package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("my_secret_key")

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return JwtKey, nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("user_id", uint(claims["user_id"].(float64)))
            c.Next()
        } else {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
        }
    }
}
