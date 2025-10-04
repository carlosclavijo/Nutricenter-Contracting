package deliveries

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewDelivery(t *testing.T) {
	id, contractId := uuid.New(), uuid.New()
	date := time.Now().AddDate(0, 0, 3)
	street, number := "123 Main St", 456
	lat, lon := 40.7128, -74.0060
	coordinates, _ := valueobjects.NewCoordinates(lat, lon)
	status, _ := ParseDeliveryStatus("pending")
	createdAt, updatedAt, deletedAt := time.Now(), time.Now(), (*time.Time)(nil)

	delivery := NewDelivery(contractId, date, street, number, coordinates)

	assert.NotNil(t, delivery)
	assert.NotNil(t, delivery.Id())
	assert.Equal(t, contractId, delivery.ContractId())
	assert.Equal(t, date, delivery.Date())
	assert.Equal(t, street, delivery.Street())
	assert.Equal(t, number, delivery.Number())
	assert.Equal(t, coordinates, delivery.Coordinates())
	assert.Equal(t, status, delivery.Status())
	assert.NotNil(t, delivery.CreatedAt())
	assert.NotNil(t, delivery.UpdatedAt())
	assert.Nil(t, delivery.DeletedAt())

	delivery = NewDeliveryFromDB(id, contractId, date, street, number, 40.7128, -74.0060, "pending", createdAt, updatedAt, deletedAt)

	assert.NotNil(t, delivery)
	assert.Equal(t, id, delivery.Id())
	assert.Equal(t, contractId, delivery.ContractId())
	assert.Equal(t, date, delivery.Date())
	assert.Equal(t, street, delivery.Street())
	assert.Equal(t, number, delivery.Number())
	w
	dCoordinates := delivery.Coordinates()
	assert.Equal(t, lat, dCoordinates.Latitude())
	assert.Equal(t, lon, dCoordinates.Longitude())
	assert.Exactly(t, coordinates, delivery.Coordinates())
	assert.Equal(t, status, delivery.Status())
	assert.Equal(t, createdAt, delivery.CreatedAt())
	assert.Equal(t, updatedAt, delivery.UpdatedAt())
	assert.Equal(t, deletedAt, delivery.DeletedAt())

	street = "456 Elm St"
	number = 789
	coordinates, _ = valueobjects.NewCoordinates(34.0522, -118.2437)

	delivery.Update(street, number, coordinates)

	assert.Equal(t, street, delivery.Street())
	assert.Equal(t, number, delivery.Number())
	assert.Equal(t, coordinates, delivery.Coordinates())
	assert.False(t, delivery.UpdatedAt().Before(updatedAt))
}
