package administratorhandlers

import (
	"context"
	administratorqueries "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries/administrator"
	"log"
)

func (h *AdministratorHandler) HandleCountAll(ctx context.Context, qry administratorqueries.CountAllAdministratorsQuery) (int, error) {
	count, err := h.repository.CountAll(ctx)
	if err != nil {
		log.Printf("[handler:administrator][HandleCountAll] error counting all administrators: %v", err)
		return 0, err
	}
	return count, nil
}
