package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/commands"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	"log"
)

func (h *ContractHandler) HandleChangeStatus(ctx context.Context, cmd commands.ChangeStatusContractCommand) (*contracts.Contract, error) {
	contract, err := h.repository.GetById(ctx, cmd.Id)
	if err != nil {
		log.Printf("[handler:contract][HandleChangeStatus] error getting contrac: %v", err)
		return nil, err
	}

	status, err := contracts.ParseContractStatus(cmd.Status)
	if err != nil {
		log.Printf("[handler:contract][HandleChangeStatus] error parsing contract status: %v", err)
		return nil, err
	}

	s := contract.ContractStatus
	if s == status.String() {
		log.Printf("[handler:contract][HandleChangeStatus] contract status is already set to %s", status.String())
		return nil, errors.New("contract status is already set to " + status.String())
	} else if s == "created" && status.String() != "active" {
		log.Printf("[handler:contract][HandleChangeStatus] contract '%s' cannot change from created to completed", status.String())
		return nil, errors.New("contract cannot change from created to completed")
	} else if s == "active" && status.String() != "completed" {
		log.Printf("[handler:contract][HandleChangeStatus] contract '%s' can change from active to created", status.String())
		return nil, errors.New("contract cannot change from active to created")
	} else if s == "completed" {
		log.Printf("[handler:contract][HandleChangeStatus] contract '%s' cannot change completed status", status.String())
		return nil, errors.New("contract cannot change completed status")
	}

	newContract, err := h.repository.ChangeStatus(ctx, cmd.Id, cmd.Status)
	if err != nil {
		log.Printf("[handler:contract][HandleChangeStatus] error changing status in Contract exists: %v", err)
		return nil, err
	}

	return newContract, nil
}
