package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"log"
)

func (h *PatientHandler) HandleCountAll(ctx context.Context, qry queries.CountAllPatientsQuery) (int, error) {
	count, err := h.repository.CountAll(ctx)
	if err != nil {
		log.Printf("[handler:patient][HandleCountAll] error counting all patients: %v", err)
		return 0, err
	}
	return count, nil
}
