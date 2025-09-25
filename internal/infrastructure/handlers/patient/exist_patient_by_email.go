package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/queries"
	"log"
)

func (h *PatientHandler) HandleExistByEmail(ctx context.Context, qry queries.ExistPatientByEmailQuery) (bool, error) {
	exist, err := h.repository.ExistByEmail(ctx, qry.Email)
	if err != nil {
		log.Printf("[handler:patient][HandleExistByEmail] error proving if an patient exist: %v", err)
		return false, err
	}
	return exist, err
}
