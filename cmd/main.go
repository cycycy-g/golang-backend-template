package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"your-project-name/config"
	"your-project-name/internal/server"
	"your-project-name/internal/store"
)

func main() {
	// Load configuration
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load configs:", err)
	}

	// Create and start server
	db, err := pgxpool.New(context.Background(), conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer db.Close()

	store := store.NewStore(db)
	s, err := server.NewServer(conf, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// Start the server
	if err := s.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
