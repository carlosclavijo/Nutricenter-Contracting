package dto

import (
	"github.com/google/uuid"
	"time"
)

type ContractDTO struct {
	Id              uuid.UUID     `json:"id"`
	AdministratorId uuid.UUID     `json:"administratorId"`
	PatientId       uuid.UUID     `json:"patientId"`
	ContractType    string        `json:"contractType"`
	ContractStatus  string        `json:"contractStatus"`
	CreationDate    time.Time     `json:"creationDate"`
	StartDate       time.Time     `json:"startDate"`
	EndDate         *time.Time    `json:"endDate,omitempty"`
	CostValue       float64       `json:"costValue"`
	Deliveries      []DeliveryDTO `json:"deliveries"`
}
