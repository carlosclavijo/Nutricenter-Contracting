package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"log"
)

func (h *AdministratorHandler) HandleGetList(ctx context.Context, qry queries.GetListAdministratorsQuery) (*[]dto.AdministratorDTO, error) {
	administrators, err := h.repository.GetList(ctx)
	if err != nil {
		log.Printf("[handler:administrator][HandleGetList] error getting administrators list: %v", err)
		return nil, err
	}
	return administrators, nil
}
