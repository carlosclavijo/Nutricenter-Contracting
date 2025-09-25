package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"log"
)

func (h *AdministratorHandler) HandleCountDeleted(ctx context.Context, qry queries.CountDeletedAdministratorsQuery) (int, error) {
	count, err := h.repository.CountDeleted(ctx)
	if err != nil {
		log.Printf("[handler:administrator][HandleCountDeleted] error counting deleted administrators: %v", err)
		return 0, err
	}
	return count, nil
}
