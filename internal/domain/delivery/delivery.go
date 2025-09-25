package deliveries

import (
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/google/uuid"
	"time"
)

type Delivery struct {
	*abstractions.Entity
	contractId uuid.UUID
	date       time.Time
	street     string
	number     int
}

func NewDelivery(contractId uuid.UUID, date time.Time, street string, number int) *Delivery {
	return &Delivery{
		Entity:     abstractions.NewEntity(uuid.New()),
		contractId: contractId,
		date:       date,
		street:     street,
		number:     number,
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

func (d *Delivery) String() string {
	return d.street
}

func (d *Delivery) Number() int {
	return d.number
}

func (d *Delivery) Update(street string, number int) {
	d.street = street
	d.number = number
}

func (d *Delivery) Delete() {

}
