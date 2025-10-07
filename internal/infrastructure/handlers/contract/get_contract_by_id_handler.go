package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/queries"
	"log"
)

func (h *ContractHandler) HandleGetById(ctx context.Context, qry queries.GetContractByIdQuery) (*dto.ContractDTO, error) {
	contract, err := h.repository.GetById(ctx, qry.Id)
	if err != nil {
		log.Printf("[handler:contract][HandleGetById] error getting contract by its id: %v", err)
		return nil, err
	}

	var deliveriesDTO []*dto.DeliveryDTO
	for _, delivery := range contract.Deliveries() {
		deliveryDTO := mappers.MapToDeliveryDTO(&delivery)
		deliveriesDTO = append(deliveriesDTO, deliveryDTO)
	}

	contractDTO := mappers.MapToContractDTO(contract)

	return contractDTO, nil
}
