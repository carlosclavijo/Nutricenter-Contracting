package mappers

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/application/contract/dto"
	deliveries "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"time"
)

func MapToDeliveryDTO(delivery *deliveries.Delivery) *dto.DeliveryDTO {
	c := delivery.Coordinates()
	return &dto.DeliveryDTO{
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

func MapToDeliveryResposnse(delivery *dto.DeliveryDTO, created, updated time.Time, deleted *time.Time) *dto.DeliveryResponse {
	return &dto.DeliveryResponse{
		DeliveryDTO: *delivery,
		CreatedAt:   created,
		UpdatedAt:   updated,
		DeletedAt:   deleted,
	}
}
