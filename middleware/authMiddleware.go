package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rahel-yab/Restuarant-Management/helpers"
)

// Authentication returns a Gin middleware that validates JWT access tokens
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		// Expect header in form "Bearer <token>"
		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// attach claims to context for handlers to use
		c.Set("user", claims)
		c.Next()
	}
}