package dto

/*import (
	administrators "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/administrator"
	deliveries "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestMapToContractDTO(t *testing.T) {
	id := uuid.New()
	administratorId := uuid.New()
	patientId := uuid.New()
	contractType := "half-month"
	contractStatus := "created"
	creationDate := time.Now()
	startDate := time.Now().AddDate(0, 0, 3)
	endDate := time.Now().AddDate(0, 0, 17)
	costValue := 1000

	street := "Elm Street"
	number := 30
	lat := -40.23
	lon := 72.83
	coordinates, _ := valueobjects.NewCoordinates(lat, lon)

	var ds []DeliveryDTO
	delivery := deliveries.NewDelivery(id, startDate, street, number, coordinates)
	deliveryDb := MapToDeliveryDTO(*delivery)
	ds = append(ds, deliveryDb)
	delivery = deliveries.NewDelivery(id, startDate.AddDate(0, 0, 1), street, number, coordinates)
	deliveryDb = MapToDeliveryDTO(*delivery)
	ds = append(ds, deliveryDb)
	delivery = deliveries.NewDelivery(id, startDate.AddDate(0, 0, 2), street, number, coordinates)
	deliveryDb = MapToDeliveryDTO(*delivery)
	ds = append(ds, deliveryDb)

	firstName := "John"
	lastName := "Doe"
	email, _ := valueobjects.NewEmail("user@email.com")
	password, _ := valueobjects.NewPassword("Abc123!!")
	gender := valueobjects.Male
	birth, _ := valueobjects.NewBirthDate(time.Now().AddDate(-20, 0, 0))
	phoneStr := "78978978"
	phone, _ := valueobjects.NewPhone(&phoneStr)
}*/
