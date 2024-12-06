package user

import (
	"github.com/google/uuid"
	"your-project-name/internal/handlers/core"
)

type UpdateProfileRequest struct {
	FullName string `json:"full_name" binding:"omitempty,min=2"`
	Email    string `json:"email" binding:"omitempty,email"`
}

type ProfileResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
}

type Handler struct {
	*core.BaseHandler
}
