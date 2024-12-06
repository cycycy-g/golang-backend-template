package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"your-project-name/config"
	"your-project-name/internal/auth"
	"your-project-name/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	config     config.Config
	Store      store.Store
	Router     *gin.Engine
	TokenMaker auth.Maker
	db         *pgxpool.Pool
}

func NewServer(config config.Config, store store.Store) (*Server, error) {
	tokenMaker, err := auth.NewJWTMaker(config.JWTSecret)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		Store:      store, // Use provided store
		Router:     gin.Default(),
		TokenMaker: tokenMaker,
	}

	server.setupRoutes()
	return server, nil
}

func (s *Server) Start() error {
	// Set server mode based on environment
	if s.config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Graceful shutdown setup
	srv := &http.Server{
		Addr:    s.config.ServerAddress,
		Handler: s.Router,
	}

	// Start server in a goroutine
	go func() {
		if err := s.Router.Run(s.config.ServerAddress); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	// Close database connection
	if s.db != nil {
		s.db.Close()
	}

	log.Println("Server exited successfully")
	return nil
}
