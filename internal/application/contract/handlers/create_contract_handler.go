package handlers

import (
	"context"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	"log"
)

func (h *ContractHandler) HandleCreate(ctx context.Context, cmd commands.CreateContractCommand) (*contracts.Contract, error) {
	cType, err := contracts.ParseContractType(cmd.ContractType)
	if err != nil {
		log.Printf("[handler:contract][HandleCreate] error parsing contract type: %v", err)
		return nil, err
	}

	contractFactory, err := h.factory.Create(cmd.AdministratorId, cmd.PatientId, cType, cmd.StartDate, cmd.Cost, cmd.Street, cmd.Number)
	if err != nil {
		log.Printf("[handler:contract][HandleCreate] error creating contract: %v", err)
		return nil, err
	}

	return contractFactory, nil
}
