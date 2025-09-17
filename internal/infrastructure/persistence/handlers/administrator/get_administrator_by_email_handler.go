package administratorhandlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"log"
)

func (h *AdministratorHandler) HandleGetByEmail(ctx context.Context, qry administratorqueries.GetAdministratorByEmailQuery) (*administrators.Administrator, error) {
	administrator, err := h.repository.GetByEmail(ctx, qry.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandlerGetByEmail] error getting administrator by its email: %v", err)
		return nil, err
	}
	return administrator, err
}
