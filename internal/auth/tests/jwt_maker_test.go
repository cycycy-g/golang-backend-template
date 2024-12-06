package tests

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"your-project-name/internal/auth"
	"your-project-name/internal/common/utils"
)

func TestJWTMaker(t *testing.T) {
	maker, err := auth.NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	userID, err := uuid.NewUUID()
	require.NoError(t, err)

	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, userID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
