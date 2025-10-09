package deliveries

import (
	"errors"
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"time"
)

var (
	ErrNotPendingDelivery         = errors.New("delivery is not pending so you can't update it")
	ErrCannotChangeDeliveryStatus = errors.New("cannot make that status change")
	ErrNotADeliveryStatus         = errors.New("not a delivery status")
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

func (d *Delivery) Update(street string, number int, coordinates valueobjects.Coordinates) error {
	if d.Status() != Pending {
		return ErrNotPendingDelivery
	}
	d.street = street
	d.number = number
	d.coordinates = coordinates
	d.updatedAt = time.Now()

	return nil

}

func (d *Delivery) ChangeStatus(status DeliveryStatus) error {
	if status != Delivered && status != Cancelled && d.status != Pending {
		return fmt.Errorf("%w: got %s", ErrCannotChangeDeliveryStatus, status)
	}

	d.status = status
	return nil
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

func NewDeliveryFromDB(id, contractId uuid.UUID, date time.Time, street string, number int, latitude, longitude float64, status string, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) (*Delivery, error) {
	coordinates, err := valueobjects.NewCoordinates(latitude, longitude)
	if err != nil {
		return nil, err
	}

	newStatus, err := ParseDeliveryStatus(status)
	if err != nil {
		return nil, err
	}

	return &Delivery{
		Entity:      abstractions.NewEntity(id),
		contractId:  contractId,
		date:        date,
		street:      street,
		number:      number,
		coordinates: coordinates,
		status:      newStatus,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		deletedAt:   deletedAt,
	}, nil
}
