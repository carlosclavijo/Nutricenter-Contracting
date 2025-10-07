package mappers

import (
	deliveries "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMapToDeliveryDTO(t *testing.T) {
	contractId := uuid.New()
	date := time.Now()
	street := "Elm Street"
	number := 30
	lat := -40.23
	lon := 72.83
	coordinates, err := valueobjects.NewCoordinates(lat, lon)
	assert.NoError(t, err)

	delivery := deliveries.NewDelivery(contractId, date, street, number, coordinates)
	dto := MapToDeliveryDTO(delivery)
	dtoCoord := delivery.Coordinates()

	assert.NotNil(t, dto)

	assert.Equal(t, delivery.Id().String(), dto.Id)
	assert.Equal(t, delivery.ContractId().String(), dto.ContractId)
	assert.Equal(t, delivery.Date(), dto.Date)
	assert.Equal(t, delivery.Street(), dto.Street)
	assert.Equal(t, delivery.Number(), dto.Number)
	assert.Equal(t, dtoCoord.Latitude(), dto.Latitude)
	assert.Equal(t, dtoCoord.Longitude(), dto.Longitude)
	assert.Equal(t, delivery.Status().String(), dto.Status)

	response := MapToDeliveryResposnse(dto, delivery.CreatedAt(), delivery.UpdatedAt(), delivery.DeletedAt())

	assert.NotNil(t, response)

	assert.Exactly(t, *dto, response.DeliveryDTO)
	assert.Equal(t, delivery.CreatedAt(), response.CreatedAt)
	assert.Equal(t, delivery.UpdatedAt(), response.UpdatedAt)
	assert.Equal(t, delivery.DeletedAt(), response.DeletedAt)

}
