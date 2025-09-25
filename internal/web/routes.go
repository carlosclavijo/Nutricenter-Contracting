package web

import (
	"database/sql"
	administrator "github.com/carlosclavijo/Nutricenter-Contracting/internal/web/controllers/administrator"
	patient "github.com/carlosclavijo/Nutricenter-Contracting/internal/web/controllers/patient"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	AdministratorController *administrator.AdministratorController
	PatientController       *patient.PatientController
}

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		AdministratorController: administrator.NewAdministratorHandler(db),
		PatientController:       patient.NewPatientHandler(db),
	}
}

func (r *Routes) Router() chi.Router {
	mux := chi.NewRouter()

	mux.Route("/administrators", r.AdministratorController.RegisterRoutes)
	mux.Route("/patients", r.PatientController.RegisterRoutes)

	return mux
}
