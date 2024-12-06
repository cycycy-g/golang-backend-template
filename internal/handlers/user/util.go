package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

func ValidateFullName(name string) error {
	// Add validation logic
	return nil
}

func getUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, fmt.Errorf("user id not found in context")
	}

	id, ok := userID.(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid user id format")
	}

	return id, nil
}
