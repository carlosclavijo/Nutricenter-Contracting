package handlers

import "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"

type PatientHandler struct {
	repository patients.PatientRepository
	factory    patients.PatientFactory
}

func NewPatientHandler(r patients.PatientRepository, f patients.PatientFactory) *PatientHandler {
	return &PatientHandler{
		repository: r,
		factory:    f,
	}
}
