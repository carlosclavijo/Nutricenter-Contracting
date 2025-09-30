package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"log"
)

func (h *ContractHandler) HandleCreate(ctx context.Context, cmd commands.CreateContractCommand) (*contracts.Contract, error) {
	cType, err := contracts.ParseContractType(cmd.ContractType)
	if err != nil {
		log.Printf("[handler:contract][HandleCreate] error parsing contract type: %v", err)
		return nil, err
	}

	coordinates, err := valueobjects.NewCoordinates(cmd.Longitude, cmd.Latitude)
	if err != nil {
		log.Printf("[handler:contract][HandleCreate] error creating coordinates: %v", err)
		return nil, err
	}

	contractFactory, err := h.factory.Create(cmd.AdministratorId, cmd.PatientId, cType, cmd.StartDate, cmd.Cost, cmd.Street, cmd.Number, *coordinates)
	if err != nil {
		log.Printf("[handler:contract][HandleCreate] error creating contract factory: %v", err)
		return nil, err
	}

	contract, err := h.repository.Create(ctx, contractFactory)
	if err != nil {
		log.Printf("[handler:contract][HandleCreate] error creating contract: %v", err)
		return nil, err
	}

	log.Printf("[handler:contract][HandleCreate] contract created")
	return contract, nil
}
