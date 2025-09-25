package main

import (
	"log"

	"github.com/mahdifr17/CheckoutService/cmd"
	"github.com/mahdifr17/CheckoutService/internals/config"
	"github.com/mahdifr17/CheckoutService/internals/database"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.Connect(cfg.Database.URL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	cmd.StartAPI(cfg, db)
}
