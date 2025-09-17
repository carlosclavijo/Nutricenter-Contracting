package web

import (
	"database/sql"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	AdministratorHandler *controllers.AdministratorController
}

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		AdministratorHandler: controllers.NewAdministratorHandler(db),
	}
}

func (r *Routes) Router() chi.Router {
	mux := chi.NewRouter()

	mux.Route("/administrators", r.AdministratorHandler.RegisterRoutes)

	return mux
}
