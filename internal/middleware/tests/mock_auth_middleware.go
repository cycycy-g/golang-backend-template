package tests

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
	"your-project-name/internal/auth"
)

func AddAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker auth.Maker,
	userID uuid.UUID,
	duration time.Duration,
) {
	token, err := tokenMaker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	authorizationHeader := fmt.Sprintf("%s %s", "bearer", token)
	request.Header.Set("Authorization", authorizationHeader)
}
