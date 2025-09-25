package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"log"
)

func (h *PatientHandler) HandleCountDeleted(ctx context.Context, qry queries.CountDeletedPatientsQuery) (int, error) {
	count, err := h.repository.CountDeleted(ctx)
	if err != nil {
		log.Printf("[handler:patient][HandleCountDeleted] error counting deleted patients: %v", err)
		return 0, err
	}
	return count, nil
}
