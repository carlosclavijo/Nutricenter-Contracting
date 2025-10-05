package dto

import (
	adm "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	ptn "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
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
