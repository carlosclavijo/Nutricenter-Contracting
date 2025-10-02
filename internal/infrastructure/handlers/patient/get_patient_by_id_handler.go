package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"log"
)

func (h *PatientHandler) HandleGetById(ctx context.Context, qry queries.GetPatientByIdQuery) (*dto.PatientDTO, error) {
	patients, err := h.repository.GetById(ctx, qry.Id)
	if err != nil {
		log.Printf("[handler:patient][HandleGetById] error getting patient by its id: %v", err)
		return nil, err
	}

	patientsDTO := dto.MapToPatientDTO(patients)
	return patientsDTO, nil
}
