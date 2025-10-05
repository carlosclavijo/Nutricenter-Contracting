package dto

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
	coordinates, _ := valueobjects.NewCoordinates(lat, lon)

	delivery := deliveries.NewDelivery(contractId, date, street, number, coordinates)
	deliveryDb := MapToDeliveryDTO(*delivery)

	assert.NotNil(t, deliveryDb)
	assert.Equal(t, deliveryDb.Id, delivery.Id().String())
	assert.Equal(t, deliveryDb.ContractId, delivery.ContractId().String())
	assert.Equal(t, deliveryDb.Date, delivery.Date())
	assert.Equal(t, deliveryDb.Street, delivery.Street())
	assert.Equal(t, deliveryDb.Number, delivery.Number())
	v := delivery.Coordinates()
	assert.Equal(t, deliveryDb.Latitude, v.Latitude())
	assert.Equal(t, deliveryDb.Longitude, v.Longitude())
	assert.Equal(t, deliveryDb.Status, delivery.Status().String())
}
