package dto

import (
	adm "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	ptn "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	contracts "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
)

func MapToContractDTO(contract contracts.Contract, deliveries []*DeliveryDTO, administrator *adm.AdministratorDTO, patient *ptn.PatientDTO) *ContractDTO {
	return &ContractDTO{
		Id:              contract.Id().String(),
		AdministratorId: contract.AdministratorId().String(),
		PatientId:       contract.PatientId().String(),
		ContractType:    contract.ContractType().String(),
		ContractStatus:  contract.ContractStatus().String(),
		CreationDate:    contract.CreationDate(),
		StartDate:       contract.StartDate(),
		EndDate:         contract.EndDate(),
		CostValue:       contract.CostValue(),
		Deliveries:      deliveries,
		Administrator:   administrator,
		Patient:         patient,
	}
}
