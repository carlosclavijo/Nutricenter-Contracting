package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"log"
)

func (h *PatientHandler) HandleGetList(ctx context.Context, qry queries.GetListPatientsQuery) (*[]dto.PatientDTO, error) {
	patients, err := h.repository.GetList(ctx)
	if err != nil {
		log.Printf("[handler:[patient]][HandleGetList] error getting patients list: %v", err)
		return nil, err
	}
	return patients, nil
}
