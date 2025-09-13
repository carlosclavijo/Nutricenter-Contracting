package main

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/persistence"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/web"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

const connection = ":8080"

func main() {
	_ = godotenv.Load("../../.env")

	db, err := persistence.NewPostgresDB()
	if err != nil {
		log.Fatalf("web/main database main error: %v", err)
		return
	}

	routes := web.NewRoutes(db)

	err = http.ListenAndServe(connection, routes.Router())
	if err != nil {
		log.Fatalf("web/main Web connection  error: %v", err)
		return
	}
}
