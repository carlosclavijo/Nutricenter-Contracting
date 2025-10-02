package handlers

import (
	"context"
	administrator "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/queries"
	patient "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"log"
)

func (h *ContractHandler) HandleGetById(ctx context.Context, qry queries.GetContractByIdQuery) (*dto.ContractDTO, error) {
	contract, err := h.repository.GetById(ctx, qry.Id)
	if err != nil {
		log.Printf("[handler:contract][HandleGetById] error getting contract by its id: %v", err)
		return nil, err
	}

	admin, err := h.repoAdmin.GetById(ctx, contract.AdministratorId())
	if err != nil {
		log.Printf("[handler:contract][HandleGetById] error getting administrator: %v", err)
		return nil, err
	}

	ptnt, err := h.repoPatient.GetById(ctx, contract.PatientId())
	if err != nil {
		log.Printf("[handler:contract][HandleGetById] error getting patient: %v", err)
		return nil, err
	}

	adminDTO := administrator.MapToAdministratorDTO(admin)
	ptntDTO := patient.MapToPatientDTO(ptnt)

	var deliveriesDTO []*dto.DeliveryDTO
	for _, delivery := range contract.Deliveries() {
		deliveryDTO := dto.MapToDeliveryDTO(&delivery)
		deliveriesDTO = append(deliveriesDTO, deliveryDTO)
	}

	contractDTO := dto.MapToContractDTO(contract, deliveriesDTO, adminDTO, ptntDTO)

	return contractDTO, nil
}
