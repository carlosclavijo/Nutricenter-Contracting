package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/google/uuid"
	"log"
)

func (h *AdministratorHandler) HandleDelete(ctx context.Context, id uuid.UUID) (*administrators.Administrator, error) {
	if id == uuid.Nil {
		log.Printf("[handler:administrator][HandleDelete] Id '%v' is nil", id)
		return nil, errors.New("the id is not valid")
	}

	exist, err := h.repository.ExistById(ctx, id)
	if err != nil {
		log.Printf("[handler:administrator][HandleDelete] error verifying if Administrator exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:administrator][HandleDelete] the Administrator doesn't exists '%v'", id)
		return nil, errors.New("administrator not found")
	}

	admin, err := h.repository.Delete(ctx, id)
	if err != nil {
		log.Printf("[handler:administrator][HandleDelete] error Deleting Administrator: %v", err)
		return nil, err
	}

	return admin, nil
}
