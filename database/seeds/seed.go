package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"user-svc/database/seeds/data"
	"user-svc/internal/shared/config"
)

func main() {
	// Define command-line flags
	dbURL := flag.String("database", "", "Database connection URL")
	flag.Parse()

	// Validate the database connection URL flag
	if *dbURL == "" {
		log.Fatal("Please provide a valid database connection URL")
	}

	// Connect to the database
	db, err := sql.Open("postgres", *dbURL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Failed to close the database connection:", err)
		}
	}(db)

	cfg := config.New()
	seeder := data.NewSeeder(db, cfg)

	// Perform data seeding
	if err = seeder.Run(); err != nil {
		log.Fatal("Failed to seeding the database:", err)
	}
	fmt.Println("Data seeding completed successfully!")
}
