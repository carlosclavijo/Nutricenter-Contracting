package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"log"
)

func (h *PatientHandler) HandleGetByEmail(ctx context.Context, qry queries.GetPatientByEmailQuery) (*dto.PatientDTO, error) {
	patient, err := h.repository.GetByEmail(ctx, qry.Email)
	if err != nil {
		log.Printf("[handler:patient][HandlerGetByEmail] error getting patient by its email: %v", err)
		return nil, err
	}

	patientDTO := mappers.MapToPatientDTO(patient)
	return patientDTO, err
}
