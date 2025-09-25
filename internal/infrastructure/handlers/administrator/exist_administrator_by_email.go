package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"log"
)

func (h *AdministratorHandler) HandleExistByEmail(ctx context.Context, qry queries.ExistAdministratorByEmailQuery) (bool, error) {
	exist, err := h.repository.ExistByEmail(ctx, qry.Email)
	if err != nil {
		log.Printf("[handler:administrator][HandleExistByEmail] error proving if an administrator exist: %v", err)
		return false, err
	}
	return exist, err
}
