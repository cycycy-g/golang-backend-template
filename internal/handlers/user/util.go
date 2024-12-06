package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

func ValidateFullName(name string) error {
	// Add validation logic
	return nil
}

func getUserIDFromContext(c *gin.Context) (int64, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, fmt.Errorf("user id not found in context")
	}

	id, ok := userID.(int64)
	if !ok {
		return 0, fmt.Errorf("invalid user id format")
	}

	return id, nil
}
