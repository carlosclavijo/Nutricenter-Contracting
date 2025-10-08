package contracts

import (
	"fmt"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/abstractions"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/delivery"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
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
	endDate         time.Time
	costValue       int
	deliveries      []deliveries.Delivery
	createdAt       time.Time
	updatedAt       time.Time
	deletedAt       *time.Time
}

func (c *Contract) Active() error {
	if c.contractStatus != Created {
		return fmt.Errorf("only Created contracts can convert to Active")
	}
	c.contractStatus = Active
	return nil
}

func (c *Contract) Completed() error {
	if c.contractStatus != Active {
		return fmt.Errorf("only Active contracts can convert to Finished")
	}
	c.contractStatus = Finished
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

func (c *Contract) EndDate() time.Time {
	return c.endDate
}

func (c *Contract) CostValue() int {
	return c.costValue
}

func (c *Contract) Deliveries() []deliveries.Delivery {
	return c.deliveries
}

func (c *Contract) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Contract) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Contract) DeletedAt() *time.Time {
	return c.deletedAt
}

func NewContract(administratorId uuid.UUID, patientId uuid.UUID, contractType ContractType, start time.Time, costValue int, street string, number int, coordinates valueobjects.Coordinates) *Contract {
	id := uuid.New()
	var endDate time.Time
	if contractType == HalfMonth {
		endDate = start.AddDate(0, 0, 14)
	} else if contractType == Monthly {
		endDate = start.AddDate(0, 0, 29)
	}
	return &Contract{
		AggregateRoot:   abstractions.NewAggregateRoot(id),
		administratorId: administratorId,
		patientId:       patientId,
		contractType:    contractType,
		contractStatus:  Created,
		creationDate:    time.Now(),
		startDate:       start,
		endDate:         endDate,
		costValue:       costValue,
		deliveries:      createCalendar(contractType, id, start, street, number, coordinates),
	}
}

func createCalendar(typ ContractType, contractId uuid.UUID, date time.Time, street string, number int, coordinates valueobjects.Coordinates) []deliveries.Delivery {
	var days []deliveries.Delivery
	if typ == HalfMonth {
		for i := 0; i < 15; i++ {
			d := deliveries.NewDelivery(contractId, date.AddDate(0, 0, 0+i), street, number, coordinates)
			days = append(days, *d)
		}
	} else if typ == Monthly {
		for i := 0; i < 30; i++ {
			d := deliveries.NewDelivery(contractId, date.AddDate(0, 0, 0+i), street, number, coordinates)
			days = append(days, *d)
		}
	}
	return days
}

func NewContractFromDb(id, aId, pId uuid.UUID, cType, cStatus string, cDate, sDate, eDate time.Time, cost int, d []deliveries.Delivery, cAt, uAt time.Time, dAt *time.Time) (*Contract, error) {
	contractType, err := ParseContractType(cType)
	if err != nil {
		return nil, fmt.Errorf("invalid contract type in DB: %w", err)
	}

	contractStatus, err := ParseContractStatus(cStatus)
	if err != nil {
		return nil, fmt.Errorf("invalid contract status in DB: %w", err)
	}

	return &Contract{
		AggregateRoot:   abstractions.NewAggregateRoot(id),
		administratorId: aId,
		patientId:       pId,
		contractType:    contractType,
		contractStatus:  contractStatus,
		creationDate:    cDate,
		startDate:       sDate,
		endDate:         eDate,
		costValue:       cost,
		deliveries:      d,
		createdAt:       cAt,
		updatedAt:       uAt,
		deletedAt:       dAt,
	}, nil
}
