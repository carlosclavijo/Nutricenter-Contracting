package administratorhandlers

import (
	"context"
	administratorqueries "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries/administrator"
	"log"
)

func (h *AdministratorHandler) HandleExistById(ctx context.Context, qry administratorqueries.ExistAdministratorByIdQuery) (bool, error) {
	exist, err := h.repository.ExistById(ctx, qry.Id)
	if err != nil {
		log.Printf("[handler:administrator][HandleExistById] error proving if an administrator exist: %v", err)
		return false, err
	}
	return exist, nil
}
