package patients

import (
	"errors"
	vo "github.com/carlosclavijo/Nutricenter-Contracting/internal/domain/valueobjects"
	"log"
	"unicode"
)

type PatientFactory interface {
	Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone vo.Phone) (*Patient, error)
}

type patientFactory struct{}

func (a patientFactory) Create(firstName, lastName string, email vo.Email, password vo.Password, gender vo.Gender, birth vo.BirthDate, phone vo.Phone) (*Patient, error) {
	if firstName == "" {
		log.Printf("[factory:patient] firstName '%s' is empty", firstName)
		return nil, errors.New("firstName is empty")
	}

	if lastName == "" {
		log.Printf("[factory:patient] lastName '%s' is empty", lastName)
		return nil, errors.New("lastName is empty")
	}

	if len(firstName) > 100 {
		log.Printf("[factory:patient] firstName '%s' is too long (length %d, maximum is 100)", firstName, len(firstName))
		return nil, errors.New("firstName is too long: maximum length is 100 characters")
	}

	if len(lastName) > 100 {
		log.Printf("[factory:patient] lastName '%s' is too long (length %d, maximum is 100)", lastName, len(lastName))
		return nil, errors.New("lastName is too long: maximum length is 100 characters")
	}

	if !isAlpha(firstName) {
		log.Printf("[factory:patient] firstName '%s' contains non-alphabetic characters", firstName)
		return nil, errors.New("firstName contains non-alphabetic characters")
	}

	if !isAlpha(lastName) {
		log.Printf("[factory:patient] lastName '%s' contains non-alphabetic characters", lastName)
		return nil, errors.New("lastName contains non-alphabetic characters")
	}

	log.Printf("[factory:patient][SUCCESS] patient created")
	return NewPatient(firstName, lastName, email, password, gender, birth, &phone), nil
}

func NewPatientFactory() PatientFactory {
	return &patientFactory{}
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) || r == ' ' {
			return false
		}
	}
	return true
}
