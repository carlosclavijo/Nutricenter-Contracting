package dto

import (
	adm "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/administrator/dto"
	ptn "github.com/carlosclavijo/Nutricenter-Contracting/internal/application/patient/dto"
	"github.com/google/uuid"
	"time"
)

type ContractDTO struct {
	Id              uuid.UUID             `json:"id"`
	AdministratorId uuid.UUID             `json:"administratorId"`
	PatientId       uuid.UUID             `json:"patientId"`
	ContractType    string                `json:"contractType"`
	ContractStatus  string                `json:"contractStatus"`
	CreationDate    time.Time             `json:"creationDate"`
	StartDate       time.Time             `json:"startDate"`
	EndDate         *time.Time            `json:"endDate,omitempty"`
	CostValue       float64               `json:"costValue"`
	Deliveries      []DeliveryDTO         `json:"deliveries"`
	Administrator   *adm.AdministratorDTO `json:"administrator"`
	Patient         *ptn.PatientDTO       `json:"patient"`
}
