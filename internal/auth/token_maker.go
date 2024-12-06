package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type Maker interface {
	CreateToken(userID uuid.UUID, duration time.Duration) (string, error)
	CreateRefreshToken(userID uuid.UUID, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
	VerifyRefreshToken(token string) (*Payload, error)
}

type Claims struct {
	UserID uuid.UUID `json:"sub"`
	Type   string    `json:"type"` // "access" or "refresh"
	jwt.RegisteredClaims
}
