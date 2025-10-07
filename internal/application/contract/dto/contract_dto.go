package dto

import (
	"time"
)

type ContractDTO struct {
	Id              string         `json:"id"`
	AdministratorId string         `json:"administratorId"`
	PatientId       string         `json:"patientId"`
	ContractType    string         `json:"contractType"`
	ContractStatus  string         `json:"contractStatus"`
	CreationDate    time.Time      `json:"creationDate"`
	StartDate       time.Time      `json:"startDate"`
	EndDate         time.Time      `json:"endDate,omitempty"`
	CostValue       int            `json:"costValue"`
	Deliveries      []*DeliveryDTO `json:"deliveries"`
}
