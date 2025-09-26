package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/google/uuid"
	"log"
)

func (h *AdministratorHandler) HandleRestore(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	if id == uuid.Nil {
		log.Printf("[handler:administrator][HandleRestore] Id '%v' is nil", id)
		return nil, errors.New("the id is not valid")
	}

	exist, err := h.repository.ExistById(ctx, id)
	if err != nil {
		log.Printf("[handler:administrator][HandleRestore] error verifying if Administrator exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:administrator][HandleRestore] the Administrator doesn't exists '%v'", id)
		return nil, errors.New("administrator not found")
	}

	admin, err := h.repository.Restore(ctx, id)
	if err != nil {
		log.Printf("[handler:administrator][HandleRestore] error Deleting Administrator: %v", err)
		return nil, err
	}

	return admin, nil
}
