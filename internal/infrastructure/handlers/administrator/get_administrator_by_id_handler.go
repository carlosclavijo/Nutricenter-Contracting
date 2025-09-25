package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"log"
)

func (h *AdministratorHandler) HandleGetById(ctx context.Context, qry queries.GetAdministratorByIdQuery) (*dto.AdministratorDTO, error) {
	administrators, err := h.repository.GetById(ctx, qry.Id)
	if err != nil {
		log.Printf("[handler:administrator][HandleGetById] error getting administrator by its id: %v", err)
		return nil, err
	}
	return administrators, nil
}
