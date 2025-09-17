package administratorhandlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries/administrator"
	"log"
)

func (h *AdministratorHandler) HandleGetAll(ctx context.Context, qry administratorqueries.GetAllAdministratorsQuery) (*[]dto.AdministratorDTO, error) {
	administrators, err := h.repository.GetAll(ctx)
	if err != nil {
		log.Printf("[handler:administrator][HandleGetAll] error getting all administrators: %v", err)
		return nil, err
	}
	return administrators, nil
}
