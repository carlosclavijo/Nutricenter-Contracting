package deliveries

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"time"
)

type Delivery struct {
	*abstractions.Entity
	contractId  uuid.UUID
	date        time.Time
	street      string
	number      int
	coordinates valueobjects.Coordinates
	status      DeliveryStatus
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

func NewDelivery(contractId uuid.UUID, date time.Time, street string, number int, coordinates valueobjects.Coordinates) *Delivery {
	return &Delivery{
		Entity:      abstractions.NewEntity(uuid.New()),
		contractId:  contractId,
		date:        date,
		street:      street,
		number:      number,
		coordinates: coordinates,
		status:      Pending,
	}
}

func NewDeliveryFromDB(id, contractId uuid.UUID, date time.Time, street string, number int, latitude, longitude float64, status string) *Delivery {
	c, _ := valueobjects.NewCoordinates(latitude, longitude)
	stt, _ := ParseDeliveryStatus(status)

	return &Delivery{
		Entity:      abstractions.NewEntity(id),
		contractId:  contractId,
		date:        date,
		street:      street,
		number:      number,
		coordinates: *c,
		status:      stt,
	}
}

func (d *Delivery) Id() uuid.UUID {
	return d.Entity.Id
}

func (d *Delivery) ContractId() uuid.UUID {
	return d.contractId
}

func (d *Delivery) Date() time.Time {
	return d.date
}

func (d *Delivery) Street() string {
	return d.street
}

func (d *Delivery) Number() int {
	return d.number
}

func (d *Delivery) Coordinates() valueobjects.Coordinates {
	return d.coordinates
}

func (d *Delivery) Status() DeliveryStatus {
	return d.status
}

func (d *Delivery) CreatedAt() time.Time {
	return d.createdAt
}

func (d *Delivery) UpdatedAt() time.Time {
	return d.updatedAt
}

func (d *Delivery) DeletedAt() *time.Time {
	return d.deletedAt
}

func (d *Delivery) Update(street string, number int) {
	d.street = street
	d.number = number
}

func (d *Delivery) Delete() {

}
