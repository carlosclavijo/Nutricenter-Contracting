package handlers

import contracts "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"

type ContractHandler struct {
	repository contracts.ContractRepository
	factory    contracts.ContractFactory
}

func NewContractHandler(r contracts.ContractRepository, f contracts.ContractFactory) *ContractHandler {
	return &ContractHandler{
		repository: r,
		factory:    f,
	}
}
