package middleware

import (
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)
var secretKey = []byte("secret_key")
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.Request.Header.Get("Authorization")
        tokenString = tokenString[len("Bearer "):]
        claims := &jwt.StandardClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return secretKey, nil
        })
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Set("userID", claims.Subject)
        c.Next()
    }
}
