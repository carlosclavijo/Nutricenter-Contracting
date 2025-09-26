package contracts

import (
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/google/uuid"
	"time"
)

type Contract struct {
	*abstractions.AggregateRoot
	administratorId uuid.UUID
	patientId       uuid.UUID
	contractType    ContractType
	contractStatus  ContractStatus
	creationDate    time.Time
	startDate       time.Time
	endDate         *time.Time
	costValue       int
	deliveries      []deliveries.Delivery
	createdAt       time.Time
	updatedAt       time.Time
	deletedAt       *time.Time
}

func NewContract(administratorId uuid.UUID, patientId uuid.UUID, contractType ContractType, start time.Time, costValue int, street string, number int) *Contract {
	id := uuid.New()
	return &Contract{
		AggregateRoot:   abstractions.NewAggregateRoot(id),
		administratorId: administratorId,
		patientId:       patientId,
		contractType:    contractType,
		contractStatus:  Created,
		creationDate:    time.Now(),
		startDate:       start,
		costValue:       costValue,
		deliveries:      createCalendar(contractType, id, start, street, number),
	}
}

func createCalendar(typ ContractType, contractId uuid.UUID, date time.Time, street string, number int) []deliveries.Delivery {
	var days []deliveries.Delivery
	if typ == HalfMonth {
		for i := 0; i < 15; i++ {
			d := deliveries.NewDelivery(contractId, date.AddDate(0, 0, 0+i), street, number)
			days = append(days, *d)
		}
	} else if typ == Monthly {
		for i := 0; i < 30; i++ {
			d := deliveries.NewDelivery(contractId, date.AddDate(0, 0, 0+i), street, number)
			days = append(days, *d)
		}
	}
	return days
}

func (c *Contract) InProgress() error {
	if c.contractStatus != Created {
		return errors.New("contract is not created")
	}
	c.contractStatus = InProgress
	return nil
}

func (c *Contract) Completed() error {
	if c.contractStatus != InProgress {
		return errors.New("contract is not in-progress")
	}
	c.contractStatus = Completed
	return nil
}

func (c *Contract) Id() uuid.UUID {
	return c.Entity.Id
}

func (c *Contract) AdministratorId() uuid.UUID {
	return c.administratorId
}

func (c *Contract) PatientId() uuid.UUID {
	return c.patientId
}

func (c *Contract) ContractType() ContractType {
	return c.contractType
}

func (c *Contract) ContractStatus() ContractStatus {
	return c.contractStatus
}

func (c *Contract) CreationDate() time.Time {
	return c.creationDate
}

func (c *Contract) StartDate() time.Time {
	return c.startDate
}

func (c *Contract) EndDate() *time.Time {
	return c.endDate
}

func (c *Contract) CostValue() int {
	return c.costValue
}

func (c *Contract) Deliveries() []deliveries.Delivery {
	return c.deliveries
}

func NewContractFromDb(id, aId, pId uuid.UUID, cType ContractType, cStatus ContractStatus, cDate, sDate time.Time, eDate *time.Time, cost int, cAt, uAt time.Time, dAt *time.Time) *Contract {
	return &Contract{
		AggregateRoot:   abstractions.NewAggregateRoot(id),
		administratorId: aId,
		patientId:       pId,
		contractType:    cType,
		contractStatus:  cStatus,
		creationDate:    cDate,
		startDate:       sDate,
		endDate:         eDate,
		costValue:       cost,
		deliveries:      []deliveries.Delivery{},
		createdAt:       cAt,
		updatedAt:       uAt,
		deletedAt:       dAt,
	}
}
