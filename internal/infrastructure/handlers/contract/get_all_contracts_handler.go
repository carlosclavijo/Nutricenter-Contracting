package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/queries"
	"log"
)

func (h *ContractHandler) HandleGetAll(ctx context.Context, qry queries.GetAllContractsQuery) (*[]dto.ContractDTO, error) {
	contracts, err := h.repository.GetAll(ctx)
	if err != nil {
		log.Printf("[handler:contract][HandleGetAll] error getting all contracts: %v", err)
		return nil, err
	}

	for i := range *contracts {
		admin, err := h.repoAdmin.GetById(ctx, (*contracts)[i].AdministratorId)
		if err != nil {
			log.Printf("[handler:contract][HandleGetAll] error getting administrator: %v", err)
			return nil, err
		}
		(*contracts)[i].Administrator = admin

		patient, err := h.repoPatient.GetById(ctx, (*contracts)[i].PatientId)
		if err != nil {
			log.Printf("[handler:contract][HandleGetAll] error getting patient: %v", err)
			return nil, err
		}
		(*contracts)[i].Patient = patient
	}

	return contracts, nil
}
