package main

import (
	"log"
	"your-project-name/config"
	"your-project-name/internal/server"
)

func main() {
	// Load configuration
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load configs:", err)
	}

	// Create and start server
	server, err := server.NewServer(conf)
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	// Start the server
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
