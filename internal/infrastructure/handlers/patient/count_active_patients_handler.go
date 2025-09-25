package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"log"
)

func (h *PatientHandler) HandleCountActive(ctx context.Context, qry queries.CountActivePatientsQuery) (int, error) {
	count, err := h.repository.CountActive(ctx)
	if err != nil {
		log.Printf("[handler:patient][HandleCountActive] error counting active patients: %v", err)
		return 0, err
	}
	return count, nil
}
