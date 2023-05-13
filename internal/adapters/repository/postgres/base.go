package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"user-svc/internal/shared/config"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) *Repository {
	const driver = "postgres"

	conn := dsn(cfg)

	db, err := sql.Open(driver, conn)

	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	return &Repository{
		db: db,
	}
}

func dsn(cfg *config.Config) string {
	host := cfg.Database.Pgsql.Host
	port := cfg.Database.Pgsql.Port
	database := cfg.Database.Pgsql.Database
	schema := cfg.Database.Pgsql.Schema
	username := cfg.Database.Pgsql.Username
	password := cfg.Database.Pgsql.Password

	connection := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	if cfg.App.Debug {
		fmt.Println(fmt.Sprintf("Trying connect postgresql with %s", connection))
	}

	return connection
}
