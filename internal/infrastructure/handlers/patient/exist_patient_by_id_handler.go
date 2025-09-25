package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"log"
)

func (h *PatientHandler) HandleExistById(ctx context.Context, qry queries.ExistPatientByIdQuery) (bool, error) {
	exist, err := h.repository.ExistById(ctx, qry.Id)
	if err != nil {
		log.Printf("[handler:patient][HandleExistById] error proving if an patient exist: %v", err)
		return false, err
	}
	return exist, nil
}
