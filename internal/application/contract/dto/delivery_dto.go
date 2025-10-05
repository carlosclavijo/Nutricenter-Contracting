package dto

import (
	"time"
)

type DeliveryDTO struct {
	Id         string    `json:"id"`
	ContractId string    `json:"contractId"`
	Date       time.Time `json:"date"`
	Street     string    `json:"street"`
	Number     int       `json:"number"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Status     string    `json:"status"`
}
