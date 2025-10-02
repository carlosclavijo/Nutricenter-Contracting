package dto

import (
	deliveries "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
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

func MapToDeliveryDTO(delivery *deliveries.Delivery) *DeliveryDTO {
	if delivery == nil {
		return nil
	}

	c := delivery.Coordinates()

	return &DeliveryDTO{
		Id:         delivery.Id().String(),
		ContractId: delivery.ContractId().String(),
		Date:       delivery.Date(),
		Street:     delivery.Street(),
		Number:     delivery.Number(),
		Latitude:   c.Latitude(),
		Longitude:  c.Longitude(),
		Status:     delivery.Status().String(),
	}
}
