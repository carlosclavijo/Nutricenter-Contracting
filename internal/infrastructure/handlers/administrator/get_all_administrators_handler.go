package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/queries"
	"log"
)

func (h *AdministratorHandler) HandleGetAll(ctx context.Context, qry queries.GetAllAdministratorsQuery) ([]*dto.AdministratorDTO, error) {
	administrators, err := h.repository.GetAll(ctx)
	if err != nil {
		log.Printf("[handler:administrator][HandleGetAll] error getting all administrators: %v", err)
		return nil, err
	}

	var adminsDTO []*dto.AdministratorDTO
	for _, admin := range administrators {
		adminDTO := dto.MapToAdministratorDTO(admin)
		adminsDTO = append(adminsDTO, adminDTO)
	}

	return adminsDTO, nil
}
