package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/mappers"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/queries"
	"log"
)

func (h *ContractHandler) HandleGetAll(ctx context.Context, qry queries.GetAllContractsQuery) ([]*dto.ContractDTO, error) {
	var contractsDTO []*dto.ContractDTO

	contracts, err := h.repository.GetAll(ctx)
	if err != nil {
		log.Printf("[handler:contract][HandleGetAll] error getting all contracts: %v", err)
		return nil, err
	}

	for i := range contracts {
		var deliveriesDTO []*dto.DeliveryDTO
		for _, delivery := range contracts[i].Deliveries() {
			deliveryDTO := mappers.MapToDeliveryDTO(&delivery)
			deliveriesDTO = append(deliveriesDTO, deliveryDTO)
		}

		contractDTO := mappers.MapToContractDTO(contracts[i])
		contractsDTO = append(contractsDTO, contractDTO)
	}

	return contractsDTO, nil
}
