package administratorhandlers

import (
	"context"
	administratorqueries "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries/administrator"
	"log"
)

func (h *AdministratorHandler) HandleCountActive(ctx context.Context, qry administratorqueries.CountActiveAdministratorsQuery) (int, error) {
	count, err := h.repository.CountActive(ctx)
	if err != nil {
		log.Printf("[handler:administrator][HandleCountActive] error counting active administrators: %v", err)
		return 0, err
	}
	return count, nil
}
