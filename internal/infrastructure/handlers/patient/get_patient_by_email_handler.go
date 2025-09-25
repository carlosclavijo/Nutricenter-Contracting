package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"log"
)

func (h *PatientHandler) HandleGetByEmail(ctx context.Context, qry queries.GetPatientByEmailQuery) (*patients.Patient, error) {
	patient, err := h.repository.GetByEmail(ctx, qry.Email)
	if err != nil {
		log.Printf("[handler:patient][HandlerGetByEmail] error getting patient by its email: %v", err)
		return nil, err
	}
	return patient, err
}
