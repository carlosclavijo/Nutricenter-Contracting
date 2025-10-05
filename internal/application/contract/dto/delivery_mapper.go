package dto

import deliveries "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"

func MapToDeliveryDTO(delivery deliveries.Delivery) DeliveryDTO {
	c := delivery.Coordinates()
	return DeliveryDTO{
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
