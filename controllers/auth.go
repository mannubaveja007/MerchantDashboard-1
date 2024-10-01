// controllers/auth.go
package controllers

import (
	"merchant-dashboard/models"
	"merchant-dashboard/utils" // Ensure this has your JWT generation logic
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User

	// Bind JSON input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	// Authenticate user (this is a placeholder; replace with your authentication logic)
	if user.Username == "username" && user.Password == "password" {
		// Generate JWT token
		token, err := utils.GenerateJWT(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
	}
}
