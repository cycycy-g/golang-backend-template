package middleware

import (
	"net/http"
	"strings"
	"your-project-name/internal/auth"

	"github.com/gin-gonic/gin"
)

// TokenMaker interface for JWT operations
type TokenMaker interface {
	VerifyToken(token string) (*TokenPayload, error)
}

// TokenPayload represents the payload of the token
type TokenPayload struct {
	UserID string
}

func AuthMiddleware(tokenMaker auth.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no authentication token provided"})
			c.Abort()
			return
		}

		payload, err := tokenMaker.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", payload.UserID)
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	// Check Authorization header
	if authHeader := c.GetHeader("Authorization"); authHeader != "" {
		fields := strings.Fields(authHeader)
		if len(fields) == 2 && strings.ToLower(fields[0]) == "bearer" {
			return fields[1]
		}
	}

	// Check cookie
	if token, err := c.Cookie("auth_token"); err == nil {
		return token
	}

	return ""
}
