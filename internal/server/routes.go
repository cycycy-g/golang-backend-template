package server

import (
	"github.com/gin-contrib/cors"
	"time"
	"your-project-name/internal/handlers"
	"your-project-name/internal/middleware"
)

func (s *Server) setupRoutes() {
	// Initialize handlers
	h := handlers.New(s.Store, s.TokenMaker)

	// Configure CORS
	s.setupCORS()

	// API routes
	api := s.Router.Group("/api")
	{
		// Auth routes (public)
		//auth := api.Group("/auth")
		//{
		//	auth.POST("/register", h.Auth.Register)
		//	auth.POST("/login", h.Auth.Login)
		//	auth.POST("/refresh", h.Auth.RefreshToken)
		//}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(s.TokenMaker))
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("/me", h.User.GetProfile)
			}

			// Add more route groups here as needed
		}
	}
}

func (s *Server) setupCORS() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins:     s.config.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
