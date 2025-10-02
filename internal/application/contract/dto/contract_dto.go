package dto

import (
	adm "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	ptn "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/contract"
	"time"
)

type ContractDTO struct {
	Id              string                `json:"id"`
	AdministratorId string                `json:"administratorId"`
	PatientId       string                `json:"patientId"`
	ContractType    string                `json:"contractType"`
	ContractStatus  string                `json:"contractStatus"`
	CreationDate    time.Time             `json:"creationDate"`
	StartDate       time.Time             `json:"startDate"`
	EndDate         time.Time             `json:"endDate,omitempty"`
	CostValue       int                   `json:"costValue"`
	Deliveries      []*DeliveryDTO        `json:"deliveries"`
	Administrator   *adm.AdministratorDTO `json:"administrator"`
	Patient         *ptn.PatientDTO       `json:"patient"`
}

func MapToContractDTO(contract *contracts.Contract, deliveries []*DeliveryDTO, administrator *adm.AdministratorDTO, patient *ptn.PatientDTO) *ContractDTO {
	if contract == nil {
		return nil
	}

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
