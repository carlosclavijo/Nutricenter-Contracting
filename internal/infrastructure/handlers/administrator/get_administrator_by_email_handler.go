package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"log"
)

func (h *AdministratorHandler) HandleGetByEmail(ctx context.Context, qry queries.GetAdministratorByEmailQuery) (*dto.AdministratorDTO, error) {
	administrator, err := h.repository.GetByEmail(ctx, qry.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandlerGetByEmail] error getting administrator by its email: %v", err)
		return nil, err
	}

	administratorDTO := mappers.MapToAdministratorDTO(administrator)

	return administratorDTO, err
}
