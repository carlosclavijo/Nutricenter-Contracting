package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func NewPostgresDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	log.Print("DSN:", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("[infrastructure:database] Connecting to PostgreDB: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Printf("[infrastructure:database] database connection established but ping failed: %v", err)
		return nil, err
	}
	return db, nil
}
