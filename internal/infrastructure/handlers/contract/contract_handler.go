package handlers

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/patient"
)

type ContractHandler struct {
	repository  contracts.ContractRepository
	repoAdmin   administrators.AdministratorRepository
	repoPatient patients.PatientRepository
	factory     contracts.ContractFactory
}

func NewContractHandler(r contracts.ContractRepository, rAdm administrators.AdministratorRepository, rPtn patients.PatientRepository, f contracts.ContractFactory) *ContractHandler {
	return &ContractHandler{
		repository:  r,
		repoAdmin:   rAdm,
		repoPatient: rPtn,
		factory:     f,
	}
}
