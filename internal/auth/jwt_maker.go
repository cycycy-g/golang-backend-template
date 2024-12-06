package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func (maker *JWTMaker) CreateToken(userID uuid.UUID, duration time.Duration) (string, error) {
	payload, err := NewPayload(userID, duration, AccessToken)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) CreateRefreshToken(userID uuid.UUID, duration time.Duration) (string, error) {
	payload, err := NewPayload(userID, duration, RefreshToken)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	return maker.verifyTokenWithType(token, AccessToken)
}

func (maker *JWTMaker) VerifyRefreshToken(token string) (*Payload, error) {
	return maker.verifyTokenWithType(token, RefreshToken)
}

func (maker *JWTMaker) verifyTokenWithType(token string, expectedType TokenType) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	// Verify token type
	if payload.TokenType != expectedType {
		return nil, ErrInvalidType
	}

	return payload, nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}
