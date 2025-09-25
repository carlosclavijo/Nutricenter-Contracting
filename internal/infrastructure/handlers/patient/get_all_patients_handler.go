package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"log"
)

func (h *PatientHandler) HandleGetAll(ctx context.Context, qry queries.GetAllPatientsQuery) (*[]dto.PatientDTO, error) {
	patients, err := h.repository.GetAll(ctx)
	if err != nil {
		log.Printf("[handler:patient][HandleGetAll] error getting all patients: %v", err)
		return nil, err
	}
	return patients, nil
}
