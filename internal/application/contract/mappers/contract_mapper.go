package mappers

import (
	administrator "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	patient "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	contracts "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	"time"
)

func MapToContractDTO(contract *contracts.Contract) *dto.ContractDTO {
	var deliveriesDTO []*dto.DeliveryDTO
	for _, k := range contract.Deliveries() {
		d := MapToDeliveryDTO(&k)
		deliveriesDTO = append(deliveriesDTO, d)
	}

	return &dto.ContractDTO{
		Id:              contract.Id().String(),
		AdministratorId: contract.AdministratorId().String(),
		PatientId:       contract.PatientId().String(),
		ContractType:    contract.ContractType().String(),
		ContractStatus:  contract.ContractStatus().String(),
		CreationDate:    contract.CreationDate(),
		StartDate:       contract.StartDate(),
		EndDate:         contract.EndDate(),
		CostValue:       contract.CostValue(),
		Deliveries:      deliveriesDTO,
	}
}

func MapToContractResponse(contract *dto.ContractDTO, adminDto *administrator.AdministratorDTO, patientDto *patient.PatientDTO, created, updated time.Time, deleted *time.Time) *dto.ContractResponse {
	return &dto.ContractResponse{
		ContractDTO:      *contract,
		AdministratorDTO: adminDto,
		PatientDTO:       patientDto,
		CreatedAt:        created,
		UpdatedAt:        updated,
		DeletedAt:        deleted,
	}
}
