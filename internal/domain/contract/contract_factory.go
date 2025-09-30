package contracts

import (
	"errors"
	"github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"github.com/google/uuid"
	"log"
	"time"
)

type ContractFactory interface {
	Create(administratorId, patientId uuid.UUID, contractType ContractType, start time.Time, cost int, street string, number int, coordinates valueobjects.Coordinates) (*Contract, error)
}

type contractFactory struct{}

func (c contractFactory) Create(administratorId, patientId uuid.UUID, contractType ContractType, start time.Time, cost int, street string, number int, coordinates valueobjects.Coordinates) (*Contract, error) {
	if administratorId == uuid.Nil {
		log.Printf("[factory:contract] administratorId '%s' is not a valid UUID", administratorId)
		return nil, errors.New("administratorId is not a valid UUID")
	}

	if patientId == uuid.Nil {
		log.Printf("[factory:contract] patientId '%s' is not a valid UUID", patientId)
		return nil, errors.New("patientId is not a valid UUID")
	}

	if contractType != HalfMonth && contractType != Monthly {
		log.Printf("[factory:contract] contractType '%s' is invalid", contractType)
		return nil, errors.New("contractType is invalid")
	}

	if !isAtLeastTwoDaysFromToday(start) {
		log.Printf("[factory:contract] startDate '%s' is before it could be", contractType)
		return nil, errors.New("startDate is not before two days after tomorrow")
	}

	if cost <= 0 {
		log.Printf("[factory:contract] cost '%v' suppose to be a positive number", contractType)
		return nil, errors.New("cost '%d' suppose to be a positive number")
	}

	if street == "" {
		log.Printf("[factory:contract] street '%s' is empty", street)
		return nil, errors.New("street is empty")
	}

	if number <= 0 {
		log.Printf("[factory:contract] number '%d' needs to be a positive number", number)
		return nil, errors.New("number needs to be a positive number")
	}

	log.Printf("[factory:contract] contractType '%s' is valid", contractType)
	return NewContract(administratorId, patientId, contractType, start, cost, street, number, coordinates), nil
}

func isAtLeastTwoDaysFromToday(date time.Time) bool {
	today := time.Now()
	todayDate := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	targetDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	daysDiff := int(targetDate.Sub(todayDate).Hours() / 24)
	return daysDiff >= 2
}

func NewContractFactory() ContractFactory {
	return &contractFactory{}
}
