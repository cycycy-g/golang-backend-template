package user

import (
	"github.com/gin-gonic/gin"
	"log"
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
		log.Printf("cant find user from context %v", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Store.GetUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ProfileResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
	})
}
