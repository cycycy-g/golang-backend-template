package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"your-project-name/config"
	"your-project-name/internal/common/utils"
	"your-project-name/internal/store"
)

func NewTestServer(t *testing.T, store store.Store) *Server {
	gin.SetMode(gin.TestMode)
	configN := config.Config{
		JWTSecret:    utils.RandomString(32),
		AllowOrigins: []string{"*"},
		Environment:  "test",
	}

	server, err := NewServer(configN, store)
	require.NoError(t, err)

	server.Store = store

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
