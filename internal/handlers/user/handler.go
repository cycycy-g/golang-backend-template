package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"your-project-name/internal/handlers/core"
)

func New(base *core.BaseHandler) *Handler {
	return &Handler{
		BaseHandler: base,
	}
}
func (h *Handler) GetProfile(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Store.GetUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, ProfileResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
	})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	// Implementation
}
