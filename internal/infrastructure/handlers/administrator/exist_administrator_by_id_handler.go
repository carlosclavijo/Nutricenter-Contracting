package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"log"
)

func (h *AdministratorHandler) HandleExistById(ctx context.Context, qry queries.ExistAdministratorByIdQuery) (bool, error) {
	exist, err := h.repository.ExistById(ctx, qry.Id)
	if err != nil {
		log.Printf("[handler:administrator][HandleExistById] error proving if an administrator exist: %v", err)
		return false, err
	}
	return exist, nil
}
