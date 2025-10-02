package handlers

import (
	"context"
	administratorDTO "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/queries"
	patientDTO "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
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
		admin, err := h.repoAdmin.GetById(ctx, (contracts)[i].AdministratorId())
		if err != nil {
			log.Printf("[handler:contract][HandleGetAll] error getting administrator: %v", err)
			return nil, err
		}

		patient, err := h.repoPatient.GetById(ctx, (contracts)[i].PatientId())
		if err != nil {
			log.Printf("[handler:contract][HandleGetAll] error getting patient: %v", err)
			return nil, err
		}

		adminDTO := administratorDTO.MapToAdministratorDTO(admin)
		ptntDTO := patientDTO.MapToPatientDTO(patient)

		var deliveriesDTO []*dto.DeliveryDTO
		for _, delivery := range contracts[i].Deliveries() {
			deliveryDTO := dto.MapToDeliveryDTO(&delivery)
			deliveriesDTO = append(deliveriesDTO, deliveryDTO)
		}

		contractDTO := dto.MapToContractDTO(contracts[i], deliveriesDTO, adminDTO, ptntDTO)
		contractsDTO = append(contractsDTO, contractDTO)
	}

	return contractsDTO, nil
}
