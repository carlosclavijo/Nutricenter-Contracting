package dto

import (
	"github.com/google/uuid"
	"time"
)

type DeliveryDTO struct {
	Id         uuid.UUID `json:"id"`
	ContractId uuid.UUID `json:"contractId"`
	Date       time.Time `json:"date"`
	Street     string    `json:"street"`
	Number     string    `json:"number"`
}
