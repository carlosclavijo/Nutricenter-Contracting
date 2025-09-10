package main

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/infrastructure/persistence"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/web"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	_ = godotenv.Load("../../.env")

	db, err := persistence.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	routes := web.NewRoutes(db)

	err = http.ListenAndServe(":8080", routes.Router())
	if err != nil {
		log.Fatal(err)
	}
}
