package web

import (
	"database/sql"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	AdministratorController *controllers.AdministratorController
	PatientController       *controllers.PatientController
	ContractController      *controllers.ContractController
}

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		AdministratorController: controllers.NewAdministratorHandler(db),
		PatientController:       controllers.NewPatientHandler(db),
		ContractController:      controllers.NewContractHandler(db),
	}
}

func (r *Routes) Router() chi.Router {
	mux := chi.NewRouter()

	mux.Route("/administrators", r.AdministratorController.RegisterRoutes)
	mux.Route("/patients", r.PatientController.RegisterRoutes)
	mux.Route("/contracts", r.ContractController.RegisterRoutes)

	return mux
}
