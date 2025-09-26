package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
	"github.com/google/uuid"
	"log"
)

func (h *PatientHandler) HandleDelete(ctx context.Context, id uuid.UUID) (*patients.Patient, error) {
	if id == uuid.Nil {
		log.Printf("[handler:patient][HandleDelete] Id '%v' is nil", id)
		return nil, errors.New("the id is not valid")
	}

	exist, err := h.repository.ExistById(ctx, id)
	if err != nil {
		log.Printf("[handler:patient][HandleDelete] error verifying if Patient exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:patient][HandleDelete] the Patient doesn't exists '%v'", id)
		return nil, errors.New("patient not found")
	}

	admin, err := h.repository.Delete(ctx, id)
	if err != nil {
		log.Printf("[handler:patient][HandleDelete] error Deleting Patient: %v", err)
		return nil, err
	}

	return admin, nil
}
