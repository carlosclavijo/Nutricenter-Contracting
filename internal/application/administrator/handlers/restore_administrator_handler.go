package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/google/uuid"
	"log"
)

func (h *AdministratorHandler) HandleRestore(ctx context.Context, id uuid.UUID) (*dto.AdministratorResponse, error) {
	if id == uuid.Nil {
		log.Printf("[handler:administrator][HandleRestore] Id '%v' is nil", id)
		return nil, administrators.ErrEmptyIdAdministrator
	}

	exist, err := h.repository.ExistById(ctx, id)
	if err != nil {
		log.Printf("[handler:administrator][HandleRestore] error verifying if Administrator exists: %v", err)
		return nil, err
	} else if !exist {
		log.Printf("[handler:administrator][HandleRestore] the Administrator doesn't exists '%v'", id)
		return nil, administrators.ErrNotFoundAdministrator
	}

	admin, err := h.repository.Restore(ctx, id)
	if err != nil {
		log.Printf("[handler:administrator][HandleRestore] error Deleting Administrator: %v", err)
		return nil, err
	}

	adminDto := mappers.MapToAdministratorDTO(admin)
	adminResponse := mappers.MapToAdministratorResponse(adminDto, admin.LastLoginAt(), admin.CreatedAt(), admin.UpdatedAt(), admin.DeletedAt())

	return adminResponse, nil
}
