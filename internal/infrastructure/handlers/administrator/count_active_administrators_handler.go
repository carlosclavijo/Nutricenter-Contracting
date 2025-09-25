package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"log"
)

func (h *AdministratorHandler) HandleCountActive(ctx context.Context, query queries.CountActiveAdministratorsQuery) (int, error) {
	count, err := h.repository.CountActive(ctx)
	if err != nil {
		log.Printf("[handler:administrator][HandleCountActive] error counting active administrators: %v", err)
		return 0, err
	}
	return count, nil
}
